package handler

import (
	"company/microservice/database"

	"github.com/gofiber/fiber/v2"
)

func GetCompany(c *fiber.Ctx) error {
	// Handle Params
	name := c.Params("name")

	res := database.GetCompany(name)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   res.Error,
		})
	}

	return c.JSON(fiber.Map{"status": "success", "mesage": "Get One Company", "data": res})
}

func CreateCompany(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "test"})
}

func UpdateCompany(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "test"})
}

func DeleteCompany(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "test"})
}
