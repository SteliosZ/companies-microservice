package handler

import "github.com/gofiber/fiber/v2"

func GetCompany(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "test"})
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
