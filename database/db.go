package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"

	"localauth/container"
	"localauth/database/migrations"
)

var db *pgx.Conn

func New(cont *container.Container) *pgx.Conn {
	if db == nil {
		var err error
		db, err = pgx.Connect(
			context.Background(),
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				cont.GetConfigData().DbHost,
				cont.GetConfigData().DbPort,
				cont.GetConfigData().DbUser,
				cont.GetConfigData().DbPassword,
				cont.GetConfigData().DbName,
				cont.GetConfigData().DbSSLMode,
			),
		)

		if err != nil {
			panic(err)
		}

		err = migrations.Run(db, context.Background())
		if err != nil {
			slog.Error(err.Error())
			panic(err)
		}
	}

	return db
}
