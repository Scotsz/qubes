package game

import (
	"context"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/config"
	"qubes/internal/model"
	"sync"
	"time"
)

type Game struct {
	cfg    *config.AppConfig
	logger *zap.SugaredLogger

	commandQueue chan *PlayerCommand

	players      *PlayerStore
	worldManager *WorldManager
	network      *NetworkManager

	tp *TickProvider
}

type PlayerCommand struct {
	id model.ClientID
	*pb.Request
}

func New(
	cfg *config.AppConfig,
	logger *zap.SugaredLogger,
	players *PlayerStore,
	worldManager *WorldManager,
	network *NetworkManager,
	tp *TickProvider,

) *Game {
	return &Game{
		cfg:          cfg,
		logger:       logger,
		commandQueue: make(chan *PlayerCommand, 20),
		players:      players,
		network:      network,
		worldManager: worldManager,
		tp:           tp,
	}
}

func (g *Game) Connect(id model.ClientID) {
	g.players.Add(id, NewPlayer())
	g.network.SendPlayerConnected(id, g.tp.Get())
}

func (g *Game) Disconnect(id model.ClientID) {
	g.players.Remove(id)
	g.network.SendPlayerDisconnected(id, g.tp.Get())
}

func (g *Game) HandleRequest(id model.ClientID, req *pb.Request) {
	g.commandQueue <- &PlayerCommand{
		id:      id,
		Request: req,
	}
}

func (g *Game) processCommands(ctx context.Context) {
	//g.logger.Info("start processing")
	//defer g.logger.Info("end processing")
	for {
		select {
		case r := <-g.commandQueue:
			g.handleCommand(ctx, g.tp.Get(), r)
		default:
			return
		}
	}
}

func (g *Game) processTickers() {
	for _, p := range g.players.All() {
		p.Tick()
	}
}

func (g *Game) Run(ctx context.Context) {
	gameTicker := time.NewTicker(time.Millisecond * 50)
	go g.worldManager.Run(ctx)
	go g.network.Run(ctx)

	defer func() {
		gameTicker.Stop()
	}()
	for {
		select {
		case <-gameTicker.C:
			{
				g.processCommands(ctx)
				g.processTickers()
				g.tp.Next()
			}
		case <-ctx.Done():
			{
				g.logger.Info("GAME STOP")
				break
			}
		}
	}
}

func (g *Game) handleCommand(ctx context.Context, tick model.TickID, cmd *PlayerCommand) {
	cmd.Request.ProtoMessage()
	switch cmd.Command.(type) {
	case *pb.Request_Move:
		g.handleMove(ctx, cmd.id, tick, cmd.GetMove())

	case *pb.Request_Shoot:
		g.handleShoot(ctx, cmd.id, tick, cmd.GetShoot())

	case *pb.Request_Changes:
		g.handleUpdatesRequest(ctx, cmd.id, tick, cmd.GetChanges())

	case *pb.Request_Connect:
		g.logger.Infof("info got connect uname:%v", cmd.GetConnect().Username)
	}
}

func (g *Game) handleMove(ctx context.Context, id model.ClientID, tick model.TickID, m *pb.Move) {
	g.logger.Infof("Got MOVE[%v:%v] ID[%v] TICK[%v]", m.Point.X, m.Point.Y, id[:8], tick)
	player, err := g.players.Get(id)
	if err != nil {
		g.logger.Info("missing player")
		return
	}
	player.SetDest(m.Point.X, m.Point.Y, m.Point.Z)
}

func (g *Game) handleShoot(ctx context.Context, id model.ClientID, tick model.TickID, m *pb.Shoot) {
	x, y, z := m.Point.X, m.Point.Y, m.Point.Z
	g.worldManager.TryRemove(Point{int(x), int(y), int(z)})
	g.logger.Infof("Got SHOOT ID[%v] TICK[%v]", id, g.tp.Get())
}

func (g *Game) handleUpdatesRequest(ctx context.Context, id model.ClientID, tick model.TickID, m *pb.UpdateRangeRequest) {

	start, end := model.TickID(m.StartTick), model.TickID(m.EndTick)
	g.logger.Info("changes requested from %v to %v", start, end)
	g.network.SendUpdateRange(ctx, id, start, end)

}

type TickProvider struct {
	mu   sync.Mutex
	tick atomic.Uint64
}

func (t *TickProvider) Get() model.TickID {
	return model.TickID(t.tick.Load())

}
func (t *TickProvider) Next() {
	t.tick.Inc()
}
