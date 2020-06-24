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
	ctx := context.Background()

	proto := protocol.NewJson()
	clientStore := ws.NewClientStore()

	sender := ws.NewSender(logger, proto, clientStore)
	g := game.New(cfg, logger, sender)

	wsServer := ws.NewServer(logger, proto, clientStore)
	wsServer.SetGame(g)

	httpServer := http.New(ctx, logger, cfg, wsServer)
	g.Run(ctx)
	log.Fatal(httpServer.Start())
}
