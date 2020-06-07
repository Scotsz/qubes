package main

import (
	"context"
	"go.uber.org/zap"
	"log"
	"qubes/internal/config"
	"qubes/internal/game"
	"qubes/internal/http"
	"qubes/internal/ws"
)

func main() {
	logger, _ := zap.NewDevelopment()
	sugared := logger.Sugar()
	appConfig, err := config.NewAppConfig("./configs/default.toml")
	if err != nil {
		log.Fatal(err)
	}

	wsServer := ws.NewServer(sugared)
	g := game.New(appConfig, wsServer, sugared)

	wsServer.SetGame(g)

	httpServer := http.New(context.Background(), sugared, appConfig, wsServer)

	go g.Run()

	log.Fatal(httpServer.Start())
}
