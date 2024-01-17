package middlewares

import "github.com/gofiber/fiber/v2"

func CheckAuth(c *fiber.Ctx) error {
	c.Next()

	return nil
}
