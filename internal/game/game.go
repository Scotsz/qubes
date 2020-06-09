package game

import (
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	pb "qubes/api"
	"qubes/internal/config"
	"qubes/internal/model"
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
	world   *World

	worldChanges  []*Change
	changeHistory map[model.TickID][]*Change
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
		world:        NewWorld(256, 256),

		worldChanges:  make([]*Change, 0),
		changeHistory: make(map[model.TickID][]*Change),
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
					m := r.Command.GetMove()
					g.logger.Infof("Got MOVE[%v:%v] ID[%v] TICK[%v]", m.Point.X, m.Point.Y, r.id[:8], r.Tick)
					g.players[r.id].SetDest(m.Point.X, m.Point.Y, m.Point.Z)
				}
			case *pb.Command_Shoot:
				{
					x, y := r.Command.GetShoot().Point.X, r.Command.GetShoot().Point.Y
					change := g.world.CalculateDestroyChange(Point{int(x), int(y), 0})
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

func (g *Game) moveChangesToHistory() {
	g.changeHistory[g.tick] = g.worldChanges
	g.worldChanges = nil
}

func (g *Game) Run() {
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
				g.processCommands()
				g.processTickers()
				g.world.ApplyChanges(g.worldChanges)
				g.tick += 1
			}
		case <-netTicker.C:
			{
				//g.logger.Infof("net tick")
				g.Broadcast(g.AllPlayersResponse())
				if len(g.worldChanges) > 0 {
					g.Broadcast(g.ChangesResponse(g.worldChanges))
					g.moveChangesToHistory()
				}
			}
		case <-g.stopCh:
			{
				g.logger.Info("GAME STOP")
				break
			}
		}
	}
}
func (g *Game) ChangesSince(tick model.TickID) []*Change {
	//TODO
	return nil
}

func (g *Game) ChangesResponse(cs []*Change) *pb.Response {
	ch := make([]*pb.Change, len(cs))
	for i, c := range cs {
		ch[i] = c.ToProto()
	}
	resp := &pb.Response{
		Tick: uint64(g.tick),
		Payload: &pb.Payload{
			Type: &pb.Payload_Changes{
				Changes: &pb.Changes{Changes: ch}}}}
	g.logger.Infof("%v", resp)
	return resp
}

func (g *Game) AllPlayersResponse() *pb.Response {
	resp := make([]*pb.Player, 0, len(g.players))
	for id, p := range g.players {
		resp = append(resp, &pb.Player{
			Id: string(id),
			Point: &pb.FloatPoint{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
			}})
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
