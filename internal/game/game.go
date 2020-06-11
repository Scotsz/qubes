package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/api"
	"qubes/internal/config"
	"qubes/internal/model"
	"qubes/internal/protocol"
	"qubes/internal/store"
	"time"
)

type Sender interface {
	Send(id model.ClientID, msg []byte)
}

type Game struct {
	cfg          *config.AppConfig
	sender       Sender
	logger       *zap.SugaredLogger
	tick         model.TickID
	commandQueue chan *PlayerCommand

	players map[model.ClientID]*Player
	world   *World

	worldChanges     []*Change
	changeRepository store.ChangeRepository
	response         *ResponseBuilder

	protocol protocol.Protocol
}
type PlayerCommand struct {
	id model.ClientID
	*pb.Request
}

func New(cfg *config.AppConfig, sender Sender, logger *zap.SugaredLogger, changerepo store.ChangeRepository) *Game {
	return &Game{
		cfg:          cfg,
		sender:       sender,
		logger:       logger,
		commandQueue: make(chan *PlayerCommand, 20),
		players:      make(map[model.ClientID]*Player),
		world:        NewWorld(256, 256, 64),

		worldChanges:     make([]*Change, 0),
		changeRepository: changerepo,
		protocol:         protocol.NewJson(),
	}

}

func (g *Game) OnConnect(id model.ClientID) {
	g.players[id] = NewPlayer()

	g.Broadcast(g.response.PlayerConnected(string(id), g.tick))
}

func (g *Game) OnDisconnect(id model.ClientID) {
	g.Broadcast(g.response.PlayerDisconnected(string(id), g.tick))
	delete(g.players, id)
}

func (g *Game) OnMessage(id model.ClientID, msg []byte) {
	req := pb.Request{}
	err := g.protocol.Unmarshal(msg, &req)

	if err != nil {
		g.logger.Info(err)
		return
	}

	g.logger.Infof("got message %T", req.Command.Type)
	g.commandQueue <- &PlayerCommand{
		id:      id,
		Request: &req,
	}
}

func (g *Game) processCommands(ctx context.Context) {
	//g.logger.Info("start processing")
	//defer g.logger.Info("end processing")
	for {
		select {
		case r := <-g.commandQueue:
			g.handleCommand(ctx, r)
		default:
			return
		}
	}
}

func (g *Game) handleCommand(ctx context.Context, c *PlayerCommand) {
	switch c.Command.Type.(type) {

	case *pb.Command_Move:
		g.handleMove(ctx, c.id, model.TickID(c.Tick), c.GetCommand().GetMove())
	case *pb.Command_Shoot:
		g.handleShoot(ctx, c.id, model.TickID(c.Tick), c.GetCommand().GetShoot())

	case *pb.Command_Changes:
		g.handleChanges(ctx, c.id, model.TickID(c.Tick), c.GetCommand().GetChanges())
	}
}

func (g *Game) handleMove(ctx context.Context, id model.ClientID, tick model.TickID, m *pb.Move) {
	g.logger.Infof("Got MOVE[%v:%v] ID[%v] TICK[%v]", m.Point.X, m.Point.Y, id[:8], tick)
	g.players[id].SetDest(m.Point.X, m.Point.Y, m.Point.Z)
}

func (g *Game) handleShoot(ctx context.Context, id model.ClientID, tick model.TickID, m *pb.Shoot) {
	x, y := m.Point.X, m.Point.Y
	change := g.world.CalculateDestroyChange(Point{int(x), int(y), 0})
	g.worldChanges = append(g.worldChanges, change)
	g.logger.Infof("Got SHOOT ID[%v] TICK[%v]", id, tick)
}

func (g *Game) handleChanges(ctx context.Context, id model.ClientID, tick model.TickID, m *pb.UpdateRangeRequest) {
	start, end := model.TickID(m.StartTick), model.TickID(m.EndTick)
	g.logger.Info("changes requested from %v to %v", start, end)
	strings, err := g.changeRepository.GetByRangeRaw(ctx, start, end)
	if err != nil {
		g.logger.Info(err)
	}
	for _, s := range strings {
		g.logger.Info(s)
		g.sender.Send(id, []byte(s))
	}
}

func (g *Game) processTickers() {
	for _, p := range g.players {
		p.Tick()
	}
}

func (g *Game) moveChangesToHistory(bytes []byte) {
	g.changeRepository.StoreRaw(context.Background(), g.tick, bytes)
	g.worldChanges = nil
}

func (g *Game) Run(ctx context.Context) {
	gameTicker := time.NewTicker(time.Millisecond * 50)
	netTicker := time.NewTicker(time.Millisecond * 100)
	defer func() {
		gameTicker.Stop()
		netTicker.Stop()
	}()
	for {
		select {
		case <-gameTicker.C:
			{
				//g.logger.Info("game tick")
				g.processCommands(ctx)
				g.processTickers()
				g.world.ApplyChanges(g.worldChanges)
				g.tick += 1
			}
		case <-netTicker.C:
			{
				//g.logger.Infof("net tick")
				g.Broadcast(g.response.AllPlayers(g.players, g.tick))
				if len(g.worldChanges) > 0 {
					changes := g.response.Changes(g.worldChanges, g.tick)
					bytes, err := g.protocol.Marshal(changes)
					if err != nil {
						g.logger.Error(err)
					}
					g.BroadcastRaw(bytes)
					g.moveChangesToHistory(bytes)
				}
			}
		case <-ctx.Done():
			{
				g.logger.Info("GAME STOP")
				break
			}
		}
	}
}

func (g *Game) BroadcastRaw(bytes []byte) {
	for i := range g.players {
		g.sender.Send(i, bytes)
	}
}

func (g *Game) Broadcast(resp *pb.Response) {
	//g.logger.Infof("broadcasting to %v", len(g.players))
	bytes, err := g.protocol.Marshal(resp)
	if err != nil {
		g.logger.Error(err)
	}
	for i := range g.players {
		g.sender.Send(i, bytes)
	}
}
