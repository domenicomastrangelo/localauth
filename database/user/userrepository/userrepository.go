package userrepository

import (
	"context"
	"errors"
	"localauth/database/user"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetUsers() (*[]user.User, error)
	GetUser(id int) (*user.User, error)
	AddUser(group *user.User, ctx context.Context) error
	EditUser(group *user.User) error
}

type RepositoryImpl struct {
	DB *pgx.Conn
}

var (
	ErrInvalidName       = errors.New("Invalid Name")
	ErrInvalidSurname    = errors.New("Invalid Surname")
	ErrInvalidPassword   = errors.New("Invalid password")
	ErrInvalidEmail      = errors.New("Invalid email")
	ErrInvalidGroupID    = errors.New("Invalid group id")
	ErrUserAlreadyExists = errors.New("User already exists")
)

func New(db *pgx.Conn) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}

func (r *RepositoryImpl) GetUsers() (*[]user.User, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT id, name, surname, email, group_id FROM users")
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	} else if err == pgx.ErrNoRows {
		return nil, nil
	}

	var users []user.User = make([]user.User, 0)
	var u = user.New()

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.GroupID)
		if err != nil {
			slog.Error(err.Error())
		}

		users = append(users, *u)
	}

	return &users, nil
}

func (r *RepositoryImpl) GetUser(id int) (*user.User, error) {
	row := r.DB.QueryRow(context.Background(), "SELECT id, name, surname, email, group_id FROM users WHERE id = $1", id)

	var u = user.New()
	err := row.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.GroupID)
	if err != nil && err != pgx.ErrNoRows {
		return &user.User{}, err
	} else if err == pgx.ErrNoRows {
		return &user.User{}, nil
	}

	return u, nil
}

func (r *RepositoryImpl) AddUser(user *user.User, ctx context.Context) error {
	_, err := r.DB.Exec(context.Background(), "INSERT INTO users (name, surname, password, email, group_id) VALUES ($1, $2, $3, $4, $5)", user.Name, user.Surname, user.Password, user.Email, user.GroupID)
	if err != nil && strings.Contains(err.Error(), "SQLSTATE 23505") {
		return ErrUserAlreadyExists
	} else if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) EditUser(user *user.User) error {
	return nil
}
