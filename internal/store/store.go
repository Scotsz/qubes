package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"qubes/internal/model"
)

type redisRepository struct {
	db *redis.Client
}

type ChangeRepository interface {
	StoreRaw(ctx context.Context, tick model.TickID, data []byte)
	GetByRangeRaw(ctx context.Context, start, end model.TickID) ([]string, error)
}

func (r redisRepository) StoreRaw(ctx context.Context, tick model.TickID, data []byte) {
	r.db.ZAdd(ctx, "updates", &redis.Z{
		Score:  float64(tick),
		Member: data,
	})
}

func (r redisRepository) GetByRangeRaw(ctx context.Context, start, end model.TickID) ([]string, error) {
	res, err := r.db.ZRangeByScore(ctx, "updates", &redis.ZRangeBy{
		Min:    fmt.Sprintf("%v", float64(start)),
		Max:    fmt.Sprintf("%v", float64(end)),
		Offset: 0,
		Count:  0,
	}).Result()

	return res, err
}

func NewChangeRepository(db *redis.Client) ChangeRepository {
	return &redisRepository{db: db}
}
