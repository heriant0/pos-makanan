package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/pos-makanan/domain/auth"
	"github.com/heriant0/pos-makanan/domain/categories"
	"github.com/heriant0/pos-makanan/external/database"
	"github.com/heriant0/pos-makanan/internal/config"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

var cfg config.Config
var postgresdb *sqlx.DB
var redisdb *redis.Client
var mongodb *mongo.Client
var err error

func init() {
	// Postgres Connection
	cfg = config.LoadConfig("config.yaml")
	postgresdb, err = database.ConnectPostgres(cfg.PostgresDB)
	if err != nil {
		panic(err)
	}

	if postgresdb != nil {
		log.Println("database connected")
	}

	// Redis Connection
	redisdb, err = database.ConnectRedis(cfg.RedisDB)
	if err != nil {
		panic(err)
	}

	if redisdb != nil {
		log.Println("redis connected")
	}
	// Mongo Connection
	mongodb, err = database.ConnectMongo(cfg.MongoDB)
	// cek apakah ada error atau engga
	if err != nil {
		panic(err)
	}

	if mongodb == nil {
		panic("mongodb not connected")
	}
	log.Println("mongodb connected")
}

func main() {

	router := fiber.New(fiber.Config{
		AppName: "POS - Makanan",
		// BodyLimit: 2 * 1024 * 1024,
		Prefork: true,
	})

	v1 := router.Group("v1")
	auth.InitRouter(v1, postgresdb)
	categories.InitRouter(v1, postgresdb)

	appPort := fmt.Sprintf(cfg.App.Port)
	err = router.Listen(appPort)
	if err != nil {
		log.Panic("cannot start the apps")
	}
}
