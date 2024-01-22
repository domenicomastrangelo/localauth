package userservice

import (
	"context"
	"errors"
	"fmt"
	"localauth/database/user"
	"localauth/database/user/userrepository"
	"log/slog"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Authorize(user *user.User) error
	ValidateUser(user *user.User) error
	AddUser(ctx context.Context, user *user.User) error
	GetUsers(ctx context.Context) (*[]user.User, error)
	GetUser(ctx context.Context, id int64) (*user.User, error)
}

type ServiceImpl struct {
	UserRepository userrepository.Repository
}

func New(userRepository userrepository.Repository) Service {
	return &ServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *ServiceImpl) Authorize(user *user.User) error {
	return nil
}

func (s *ServiceImpl) ValidateUser(user *user.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return userrepository.ErrInvalidName
	}
	if strings.TrimSpace(user.Surname) == "" {
		return userrepository.ErrInvalidSurname
	}
	if strings.TrimSpace(user.Password) == "" ||
		len(user.Password) < 12 {
		return userrepository.ErrInvalidPassword
	}
	if strings.TrimSpace(user.Email) == "" ||
		regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`).MatchString(user.Email) == false {
		return userrepository.ErrInvalidEmail
	}

	return nil
}

func (s *ServiceImpl) HashPassword(user *user.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)

	return nil
}

func (s *ServiceImpl) CheckPassword(user *user.User, password []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) AddUser(ctx context.Context, user *user.User) error {
	if err := s.ValidateUser(user); err != nil {
		return err
	}
	if err := s.HashPassword(user); err != nil {
		return err
	}

	err := s.UserRepository.AddUser(user, ctx)
	if err != nil {
		return err
	} else if user.ID == 0 {
		slog.Error("Could not add user")
		slog.Error(fmt.Sprintf("%v", user))
		return errors.New("Could not add user")
	}

	return nil
}

func (s *ServiceImpl) GetUsers(ctx context.Context) (*[]user.User, error) {
	return s.UserRepository.GetUsers()
}

func (s *ServiceImpl) GetUser(ctx context.Context, id int64) (*user.User, error) {
	return s.UserRepository.GetUser(id)
}
