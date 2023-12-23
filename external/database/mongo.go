package database

import (
	"context"
	"fmt"

	"github.com/heriant0/pos-makanan/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg config.MongoConfig) (*mongo.Client, error) {
	uri := fmt.Sprintf("%s://%s:%s@%s:%s", cfg.Driver, cfg.UserName, cfg.Password, cfg.Host, cfg.Port)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
