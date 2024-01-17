package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"localauth/config"
)

var db *pgx.Conn

func New(config *config.Config) *pgx.Conn {
	if db == nil {
		var err error
		db, err = pgx.Connect(
			context.Background(),
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				config.DbHost,
				config.DbPort,
				config.DbUser,
				config.DbPassword,
				config.DbName,
				config.DbSSLMode,
			),
		)
		if err != nil {
			panic(err)
		}
	}

	return db
}
