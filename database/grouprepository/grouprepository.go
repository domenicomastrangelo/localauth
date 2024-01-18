package grouprepository

import (
	"context"
	"errors"
	"localauth/database/group"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5"
)

var ErrInvalidGroupName = errors.New("Invalid group name")
var ErrGroupAlreadyExists = errors.New("Group already exists")

type Repository interface {
	GetGroups(ctx context.Context) (*[]group.Group, error)
	GetGroup(ctx context.Context, id int) (*group.Group, error)
	AddGroup(group *group.Group, ctx context.Context) error
	EditGroup(group *group.Group) error
}

type RepositoryImpl struct {
	DB *pgx.Conn
}

func New(db *pgx.Conn) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}

func (r *RepositoryImpl) GetGroups(ctx context.Context) (*[]group.Group, error) {
	rows, err := r.DB.Query(ctx, "SELECT id, name FROM groups")
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	} else if err == pgx.ErrNoRows {
		return nil, nil
	}

	var groups []group.Group = make([]group.Group, 0)
	var gg = group.New()

	for rows.Next() {
		err = rows.Scan(&gg.ID, &gg.Name)
		if err != nil {
			slog.Error(err.Error())
		}

		groups = append(groups, *gg)
	}

	return &groups, nil
}

func (r *RepositoryImpl) GetGroup(ctx context.Context, id int) (*group.Group, error) {
	row := r.DB.QueryRow(ctx, "SELECT id, name FROM groups WHERE id = $1", id)

	var gg = group.New()
	err := row.Scan(&gg.ID, &gg.Name)
	if err != nil && err != pgx.ErrNoRows {
		return &group.Group{}, err
	} else if err == pgx.ErrNoRows {
		return &group.Group{}, nil
	}

	return gg, nil
}

func (r *RepositoryImpl) AddGroup(group *group.Group, ctx context.Context) error {
	_, err := r.DB.Exec(ctx, "INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil && strings.Contains(err.Error(), "SQLSTATE 23505") {
		return ErrGroupAlreadyExists
	}

	return err
}

func (r *RepositoryImpl) EditGroup(group *group.Group) error {
	return nil
}
