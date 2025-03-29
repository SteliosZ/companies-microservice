package handler

import "github.com/gofiber/fiber/v2"

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "World!"})
}
