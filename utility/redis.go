package utility

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/heriant0/pos-makanan/external/database"
	"github.com/heriant0/pos-makanan/internal/config"
)

var cfgDB config.Config
var redisdb *redis.Client

func Redis(timeout time.Duration) {
	_, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()
	redisdb, err := database.ConnectRedis(cfgDB.RedisDB)
	if err != nil {
		log.Println("db not connected with error", err.Error())
		return
	}

	if redisdb == nil {
		log.Println("redis not connected with unknown error")
		return
	}
}
