package handlers

import (
	"context"
	"localauth/container"
	"localauth/database/user"
	"localauth/database/user/userrepository"
	"localauth/database/user/userservice"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsersHandler(c *fiber.Ctx, cont *container.Container) error {
	userRepository := cont.GetUserRepository()

	usersvc := userservice.New(userRepository)
	users, err := usersvc.GetUsers(context.Background())
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.JSON(users)

	return nil
}

func GetUserHandler(c *fiber.Ctx, cont *container.Container) error {
	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusBadRequest)

		return err
	}

	userRepository := cont.GetUserRepository()

	usersvc := userservice.New(userRepository)
	user, err := usersvc.GetUser(context.Background(), int64(id))
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.JSON(user)

	return nil
}

func AddUserHandler(c *fiber.Ctx, cont *container.Container) error {
	u := user.New()
	err := c.BodyParser(u)
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusBadRequest)

		return err
	}

	userRepository := cont.GetUserRepository()

	usersvc := userservice.New(userRepository)
	err = usersvc.AddUser(context.Background(), u)
	if err != nil && err == userrepository.ErrUserAlreadyExists {
		slog.Error(err.Error())
		c.Status(fiber.StatusConflict)

		return err
	} else if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusBadRequest)

		return err
	}

	c.Status(fiber.StatusCreated)

	return nil
}

func EditUserHandler(c *fiber.Ctx, cont *container.Container) error {
	return nil
}
