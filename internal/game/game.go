package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/config"
	"qubes/internal/model"
	"time"
)

type Game struct {
	worldManager   WorldManager
	requestHandler *RequestHandler
	network        *NetworkManager
	cfg            *config.AppConfig
}

func New(cfg *config.AppConfig, logger *zap.SugaredLogger, sender Sender) *Game {

	players := NewPlayerStore()
	network := NewNetworkManager(sender, logger, players)
	worldManager := NewWorldManager(logger, GetTestWorld(10), network)
	requestHandler := NewRequestHandler(logger, worldManager, network)

	return &Game{
		worldManager:   worldManager,
		requestHandler: requestHandler,
		network:        network,
		cfg:            cfg,
	}
}

func simTick(rate time.Duration) chan model.TickID {
	tick := model.TickID(0)
	timer := time.NewTicker(rate)
	ch := make(chan model.TickID)
	go func() {
		for {
			select {
			case <-timer.C:
				tick++
				ch <- tick
			}
		}
	}()
	return ch
}

func (g *Game) Connect(id model.ClientID) {
	g.worldManager.AddTicker(id, NewPlayer())
	g.network.SendPlayerConnected(id)
}

func (g *Game) Disconnect(id model.ClientID) {
	g.worldManager.RemoveTicker(id)
	g.network.SendPlayerDisconnected(id)
}

func (g *Game) HandleRequest(id model.ClientID, req *pb.Request) {
	g.requestHandler.AddRequest(id, req)
}

func (g *Game) Run(ctx context.Context) {
	gameTicker := simTick(time.Millisecond * 50) //(time.Second * time.Duration(1/g.cfg.Game.SimulationRate))

	go g.worldManager.Run(ctx, gameTicker)
	go g.requestHandler.Run(ctx)
	g.network.Run(ctx)

}
