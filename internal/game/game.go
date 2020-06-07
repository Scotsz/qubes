package game

import (
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	pb "qubes/api"
	"qubes/internal/config"
	"qubes/internal/model"
	"qubes/internal/world"
	"time"
)

type Sender interface {
	Send(id model.ClientID, response proto.Message)
}

type Game struct {
	cfg          *config.AppConfig
	sender       Sender
	logger       *zap.SugaredLogger
	tick         model.TickID
	stopCh       chan struct{}
	commandQueue chan *PlayerCommand

	players map[model.ClientID]*Player
	world   *world.World

	worldChanges  []*world.Change
	changeHistory map[model.TickID]*world.Change
}
type PlayerCommand struct {
	id model.ClientID
	*pb.Request
}

func New(cfg *config.AppConfig, sender Sender, logger *zap.SugaredLogger) *Game {
	return &Game{
		cfg:          cfg,
		sender:       sender,
		logger:       logger,
		commandQueue: make(chan *PlayerCommand, 20),
		players:      make(map[model.ClientID]*Player),

		worldChanges: make([]*world.Change, 0),
	}
}

func (g *Game) OnConnect(id model.ClientID) {
	g.players[id] = NewPlayer()

	g.Broadcast(&pb.Response{
		Tick: uint64(g.tick),
		Payload: &pb.Payload{Type: &pb.Payload_PlayerConnect{PlayerConnect: &pb.Player{
			Id: string(id),
		}}},
	})
}

func (g *Game) OnDisconnect(id model.ClientID) {
	g.Broadcast(&pb.Response{
		Tick: uint64(g.tick),
		Payload: &pb.Payload{Type: &pb.Payload_PlayerDisconnect{PlayerDisconnect: &pb.Player{
			Id: string(id),
		}}},
	})
	delete(g.players, id)
}

func (g *Game) OnMessage(id model.ClientID, req *pb.Request) {
	g.logger.Infof("got message %T", req.Command.Type)
	g.commandQueue <- &PlayerCommand{
		id:      id,
		Request: req,
	}
}

func (g *Game) processCommands() {
	//g.logger.Info("start processing")
	//defer g.logger.Info("end processing")
	for {
		select {
		case r := <-g.commandQueue:
			switch r.Command.Type.(type) {
			case *pb.Command_Move:
				{
					move := r.Command.GetMove()
					g.logger.Infof("Got MOVE[%v:%v] ID[%v] TICK[%v]", move.X, move.Y, r.id[:8], r.Tick)
					g.players[r.id].SetDest(move.X, move.Y)
				}
			case *pb.Command_Shoot:
				{
					x, y := r.Command.GetShoot().X, r.Command.GetShoot().Y
					change := g.world.CalculateDestroyChange(world.Point{int(x), int(y), 0})
					g.worldChanges = append(g.worldChanges, change)
					g.logger.Infof("Got SHOOT ID[%v] TICK[%v]", r.id, r.Tick)
				}
			}
		default:
			return
		}
	}
}

func (g *Game) processTickers() {
	for _, p := range g.players {
		p.Tick()
	}
}

func (g *Game) Run() {
	gameTicker := time.NewTicker(time.Millisecond * 25)
	netTicker := time.NewTicker(time.Millisecond * 50)
	defer func() {
		gameTicker.Stop()
		netTicker.Stop()
	}()
	for {
		select {
		case <-gameTicker.C:
			{
				//g.logger.Info("game tick")
				g.processCommands()
				g.processTickers()
				g.world.ApplyChanges(g.worldChanges)

				g.tick += 1
			}
		case <-netTicker.C:
			{
				//g.logger.Infof("net tick")
				g.Broadcast(g.AllPlayersResponse())
			}
		case <-g.stopCh:
			{
				g.logger.Info("GAME STOP")
				break
			}
		}
	}
}
func (g *Game) AllPlayersResponse() *pb.Response {
	resp := make([]*pb.Player, 0, len(g.players))
	for id, p := range g.players {
		resp = append(resp, &pb.Player{
			Id: string(id),
			X:  float32(p.X),
			Y:  float32(p.Y),
		})
	}

	r := &pb.Response{
		Tick: uint64(g.tick),
		Payload: &pb.Payload{
			Type: &pb.Payload_Players{
				Players: &pb.AllPlayers{Player: resp}},
		}}
	return r
}

func (g *Game) Stop() {
	g.stopCh <- struct{}{}
}

func (g *Game) Broadcast(msg *pb.Response) {
	//g.logger.Infof("broadcasting to %v", len(g.players))
	for i := range g.players {
		g.sender.Send(i, msg)
	}
}
