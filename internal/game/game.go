package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/config"
	"qubes/internal/model"
)

type Game struct {
	worldManager   *state
	requestHandler *RequestHandler
	network        *NetworkManager
	cfg            *config.AppConfig
}

func New(cfg *config.AppConfig, logger *zap.SugaredLogger, sender Sender) *Game {
	network := NewNetworkManager(sender, logger)
	worldManager := NewWorldManager(logger, GetTestWorld(0), network)
	requestHandler := NewRequestHandler(logger, worldManager, network)

	return &Game{
		worldManager:   worldManager,
		requestHandler: requestHandler,
		network:        network,
		cfg:            cfg,
	}
}

func (g *Game) Connect(id model.ClientID) {
	g.worldManager.AddPlayer(id)
	g.network.SendPlayerConnected(id)
}

func (g *Game) Disconnect(id model.ClientID) {
	g.worldManager.RemovePlayer(id)
	g.network.SendPlayerDisconnected(id)
}

func (g *Game) HandleRequest(id model.ClientID, req *pb.Request) {
	g.requestHandler.AddRequest(id, req)
}

func (g *Game) Run(ctx context.Context) {
	go g.worldManager.Run(ctx)
	go g.requestHandler.Run(ctx)
	go g.network.Run(ctx)

}
