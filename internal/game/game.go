package game

import (
	"context"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/config"
	"qubes/internal/model"
	"qubes/internal/protocol"
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
	protocol     protocol.Protocol

	handler CommandHandler

	tp *TickProvider
}

type PlayerCommand struct {
	id model.ClientID
	*pb.Request
}

func New(
	cfg *config.AppConfig,
	logger *zap.SugaredLogger,
	proto protocol.Protocol,
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
		protocol:     proto,
		tp:           tp,
	}
}

func (g *Game) OnConnect(id model.ClientID) {
	g.players.Add(id, NewPlayer())
	g.network.SendPlayerConnected(id, g.tp.Get())
}

func (g *Game) OnDisconnect(id model.ClientID) {
	g.players.Remove(id)
	g.network.SendPlayerDisconnected(id, g.tp.Get())
}

func (g *Game) OnMessage(id model.ClientID, msg []byte) {
	req := pb.Request{}
	err := g.protocol.Unmarshal(msg, &req) //TODO move to another goroutine

	if err != nil {
		g.logger.Info(err)
		return
	}

	g.logger.Infof("got message %T", req.Command)
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
			g.handler.Handle(ctx, g.tp.Get(), r)
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
				//g.logger.Info("game tp")
				g.processCommands(ctx)
				g.processTickers()
				//g.worldManager.ApplyUpdates(g.worldUpdates) //TODO
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
