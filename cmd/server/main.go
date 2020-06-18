package main

import (
	"context"
	"go.uber.org/zap"
	"log"
	"qubes/internal/config"
	"qubes/internal/game"
	"qubes/internal/http"
	"qubes/internal/protocol"
	"qubes/internal/ws"
)

func main() {
	lg, _ := zap.NewDevelopment()
	logger := lg.Sugar()
	cfg, err := config.NewAppConfig("./configs/default.toml")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info(cfg.Game.SimulationRate)
	ctx := context.Background()

	proto := protocol.NewJson()

	wsServer := ws.NewServer(logger, proto)

	g := game.New(cfg, logger, wsServer)

	wsServer.SetGame(g)

	httpServer := http.New(ctx, logger, cfg, wsServer)
	go g.Run(ctx)
	log.Fatal(httpServer.Start())
}
