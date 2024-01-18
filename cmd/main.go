package main

import (
	"context"
	"localauth/config"
	"localauth/container"
	"localauth/database"
	"localauth/database/grouprepository"
	"localauth/database/migrations"
	"localauth/database/user/userrepository"
	"localauth/handlers"
	"localauth/middlewares"
	"log/slog"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func main() {
	conf := config.New()

	cont := container.New()
	cont.AddElement("db", container.Element{
		Element: database.New(conf),
	})
	cont.AddElement("grouprepository", container.Element{
		Element: grouprepository.New(database.New(conf)),
	})
	cont.AddElement("userrepository", container.Element{
		Element: userrepository.New(database.New(conf)),
	})
	cont.AddElement("conf", container.Element{
		Element: conf,
	})

	err := migrations.Run(cont, context.Background())
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: conf.CookieKey,
	}))
	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-CSRF-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Strict",
	// 	Expiration:     1 * time.Hour,
	// 	KeyGenerator:   utils.UUIDv4,
	// }))

	authorizedRoutes := app.Group("/api").Use(middlewares.CheckAuth)

	authorizedRoutes.Get("/groups", func(c *fiber.Ctx) error {
		return handlers.GetGroupsHandler(c, cont)
	})

	authorizedRoutes.Get("/groups/:id", func(c *fiber.Ctx) error {
		return handlers.GetGroupHandler(c, cont)
	})

	authorizedRoutes.Post("/groups", func(c *fiber.Ctx) error {
		return handlers.AddGroupHandler(c, cont)
	})

	authorizedRoutes.Put("/groups/:id", func(c *fiber.Ctx) error {
		return handlers.EditGroupHandler(c, cont)
	})

	authorizedRoutes.Get("/users", func(c *fiber.Ctx) error {
		return handlers.GetUsersHandler(c, cont)
	})

	authorizedRoutes.Get("/users/:id", func(c *fiber.Ctx) error {
		return handlers.GetUserHandler(c, cont)
	})

	authorizedRoutes.Post("/users", func(c *fiber.Ctx) error {
		return handlers.AddUserHandler(c, cont)
	})

	authorizedRoutes.Put("/users/:id", func(c *fiber.Ctx) error {
		return handlers.EditUserHandler(c, cont)
	})

	authorizedRoutes.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000").Error())
}
