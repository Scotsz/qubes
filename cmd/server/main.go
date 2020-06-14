package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"log"
	"qubes/internal/config"
	"qubes/internal/game"
	"qubes/internal/http"
	"qubes/internal/protocol"
	"qubes/internal/store"
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
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.URL,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})
	redisClient.FlushDB(ctx)
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		logger.Fatal(err)
	}

	worldUpdateRepo := store.NewWorldUpdateRepository(redisClient)
	playerStore := game.NewPlayerStore()
	tickProvider := &game.TickProvider{}

	proto := protocol.NewJson()

	wsServer := ws.NewServer(logger, proto)
	network := game.NewNetworkManager(worldUpdateRepo, wsServer, logger, proto, playerStore, tickProvider)
	worldManager := game.NewWorldManager(logger, network).WithTestWorld()

	g := game.New(cfg, logger, playerStore, worldManager, network, tickProvider)
	wsServer.SetGame(g)
	httpServer := http.New(ctx, logger, cfg, wsServer)

	go g.Run(ctx)

	log.Fatal(httpServer.Start())
}
