package database

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/heriant0/pos-makanan/internal/config"
)

func ConnectRedis(cfg config.RedisConfig) (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})

	err := rdb.Ping(ctx)
	if err.Err() != nil {
		return nil, err.Err()
	}

	return rdb, nil
}
