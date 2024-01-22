package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Run(db *pgx.Conn, ctx context.Context) error {
	_, err := db.Exec(ctx, CreateGroupsTable())
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, CreateUsersTable())
	if err != nil {
		return err
	}

	return nil
}

func CreateGroupsTable() string {
	return `CREATE TABLE IF NOT EXISTS groups (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);

	INSERT INTO groups  (id, name) VALUES (0, 'nogroup') ON CONFLICT DO NOTHING;
	`
}

func CreateUsersTable() string {
	return `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		surname TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		group_id INTEGER REFERENCES groups(id)
	);
	`
}
