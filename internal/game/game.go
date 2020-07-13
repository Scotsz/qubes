package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/config"
	"qubes/internal/model"
)

type Game struct {
	worldManager   *stateManager
	requestHandler *RequestHandler
	network        *NetworkManager
	cfg            *config.AppConfig
	commandFactory *CommandFactory
}

func New(cfg *config.AppConfig, logger *zap.SugaredLogger, sender Sender) *Game {
	state := NewState(model.GetTestWorld(0))
	network := NewNetworkManager(sender, logger)
	cf := NewCommandFactory(network, state)
	stateManager := NewStateManager(logger, cf)
	requestHandler := NewRequestHandler(logger, stateManager, network)

	return &Game{
		worldManager:   stateManager,
		requestHandler: requestHandler,
		network:        network,
		cfg:            cfg,
		commandFactory: cf,
	}
}

func (g *Game) Connect(id model.ClientID) {
	g.worldManager.AddCommand(g.commandFactory.AddPlayer(model.PlayerID("player_" + id[:8])))
}

func (g *Game) Disconnect(id model.ClientID) {
	g.worldManager.AddCommand(g.commandFactory.RemovePlayer(model.PlayerID("player_" + id[:8])))
}

func (g *Game) HandleRequest(id model.ClientID, req *pb.Request) {
	g.requestHandler.AddRequest(id, req)
}

func (g *Game) Run(ctx context.Context) {
	go g.worldManager.Run(ctx)
	go g.requestHandler.Run(ctx)
	go g.network.Run(ctx)

}
