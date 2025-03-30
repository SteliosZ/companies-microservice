package handler

import (
	"company/microservice/database"
	"company/microservice/model"

	"github.com/gofiber/fiber/v2"
)

func GetCompany(c *fiber.Ctx) error {
	// Handle Params
	name := c.Params("name")

	err, res := database.GetCompany(name)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "object not found",
			"data":    err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"mesage": "got one company",
		"data":   res,
	})
}

func CreateCompany(c *fiber.Ctx) error {
	company := model.Company{}

	// Parse and Check Body
	if err := c.BodyParser(&company); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong input data",
			"data":    err,
		})
	}

	// db Query for Creation
	err := database.CreateCompany(company)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has created",
		"data":    company,
	})
}

func UpdateCompany(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"mesage": "got one company",
		"data":   nil,
	})
}

func DeleteCompany(c *fiber.Ctx) error {
	// Handle Params
	name := c.Params("name")

	err := database.DeleteCompany(name)
	if err != nil {
		return c.Status(409).JSON(fiber.Map{
			"status":  "error",
			"message": "object not deleted",
			"data":    err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"mesage": "deleted",
		"data":   nil,
	})
}
