package database

import (
	"fmt"

	"github.com/heriant0/pos-makanan/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {
	// create db connection
	databaseUrl := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Port, cfg.Database)
	dbConn, err := sqlx.Open(cfg.Driver, databaseUrl)

	if err != nil {
		errMessage := fmt.Errorf("error database connect: %w", err)
		panic(errMessage)
	}

	err = dbConn.Ping()
	if err != nil {
		errMessage := fmt.Errorf("error database ping: %w", err)
		panic(errMessage)
	}

	return dbConn, nil

}
