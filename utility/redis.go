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

	// 	err = client.Set(ctx, "token-user01", "ini-token-user", 10*time.Second).Err()
	// 	if err != nil {
	// 		log.Println("error when try to set data to redis with message :", err.Error())
	// 		return
	// 	}

	// 	cmd := client.Get(ctx, "token-user02")

	// res, err := cmd.Result()
	//
	//	if err != nil {
	//		log.Println("error when try to get data from redis with message :", err.Error())
	//		return
	//	}
}

// func SetData(key, value string, ttl time.Duration) error {
// 	err := redisdb.Set(context.Background(), key, value, 300*time.Second).Err()
// 	fmt.Println("🚀 ~ file: redis.go ~ line 12 ~ funcSetData ~ err : ", err)

// 	return err
// }

// func GetData(key string, redis *redis.Client) (interface{}, error) {
// 	data, err := redis.Get(context.Background(), key).Result()

// 	return data, err
// }
