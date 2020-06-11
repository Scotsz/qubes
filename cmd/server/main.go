package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"log"
	"qubes/internal/config"
	"qubes/internal/game"
	"qubes/internal/http"
	"qubes/internal/store"
	"qubes/internal/ws"
)

func main() {
	logger, _ := zap.NewDevelopment()
	sugared := logger.Sugar()
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
		sugared.Fatal(err)
	}

	changerepo := store.NewChangeRepository(redisClient)
	wsServer := ws.NewServer(sugared)
	g := game.New(cfg, wsServer, sugared, changerepo)

	wsServer.SetGame(g)

	httpServer := http.New(ctx, sugared, cfg, wsServer)

	go g.Run(ctx)

	log.Fatal(httpServer.Start())
}
