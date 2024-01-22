package main

import (
	"localauth/config"
	"localauth/container"
	"localauth/database"
	"localauth/database/grouprepository"
	"localauth/database/user/userrepository"
	"localauth/handlers"
	"localauth/middlewares"
	"os"
	"strings"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

var (
	cont *container.Container
	conf *config.ConfigData
)

func init() {
	env := os.Getenv("LOCALAUTH_ENV")
	if strings.TrimSpace(env) == "" {
		env = config.ENV_TEST
	}

	conf = config.New(env)

	cont = container.New()
	db := database.New(cont)

	cont.AddElement("db", container.Element{
		Element: db,
	})
	cont.AddElement("grouprepository", container.Element{
		Element: grouprepository.New(db),
	})
	cont.AddElement("userrepository", container.Element{
		Element: userrepository.New(db),
	})
	cont.AddElement("conf", container.Element{
		Element: conf,
	})
}

func main() {
	app := fiber.New()

	setMiddlewares(app)
	setRoutes(app)

	log.Fatal(app.Listen(":3000").Error())
}

func setMiddlewares(app *fiber.App) {
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
}

func setRoutes(app *fiber.App) {
	setAuthorizedRoutes(app)
	setUnauthorizedRoutes(app)
}

func setAuthorizedRoutes(app *fiber.App) {
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
}

func setUnauthorizedRoutes(app *fiber.App) {
}
