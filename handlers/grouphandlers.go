package handlers

import (
	"context"
	"localauth/container"
	"localauth/database/group"
	"localauth/database/grouprepository"
	"localauth/database/groupservice"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetGroupsHandler(c *fiber.Ctx, cont *container.Container) error {
	groupRepository := cont.GetGroupRepository()

	groupsvc := groupservice.New(groupRepository)
	groups, err := groupsvc.GetGroups(context.Background())
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.JSON(groups)

	return nil
}

func GetGroupHandler(c *fiber.Ctx, cont *container.Container) error {
	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusBadRequest)

		return err
	}

	groupRepository := cont.GetGroupRepository()

	groupsvc := groupservice.New(groupRepository)
	group, err := groupsvc.GetGroup(context.Background(), id)
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.JSON(group)

	return nil
}

func AddGroupHandler(c *fiber.Ctx, cont *container.Container) error {
	g := group.New()
	err := c.BodyParser(g)
	if err != nil {
		slog.Error(err.Error())
		c.Status(fiber.StatusBadRequest)

		return err
	}

	groupRepository := cont.GetGroupRepository()

	groupsvc := groupservice.New(groupRepository)
	err = groupsvc.AddGroup(context.Background(), g)

	if err != nil && err != grouprepository.ErrGroupAlreadyExists {
		slog.Error(err.Error())
		c.Status(fiber.StatusInternalServerError)

		return err
	} else if err == grouprepository.ErrGroupAlreadyExists {
		c.Status(fiber.StatusConflict)

		return err
	}

	c.Status(fiber.StatusCreated)

	return nil
}

func EditGroupHandler(c *fiber.Ctx, cont *container.Container) error {
	return nil
}
