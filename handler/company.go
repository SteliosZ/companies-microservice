package handler

import (
	"company/microservice/database"
	"company/microservice/model"

	"github.com/gofiber/fiber/v2"
)

func GetCompany(c *fiber.Ctx) error {
	// Handle Params
	companyName := c.Params("name")

	err, res := database.GetCompany(companyName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "company not found",
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
			"message": "company not created",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "company create",
		"data":    company,
	})
}

func UpdateCompany(c *fiber.Ctx) error {
	newDetails := map[string]interface{}{}

	companyName := c.Params("name")

	// Parse and Check Body
	if err := c.BodyParser(&newDetails); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong input data",
			"data":    err,
		})
	}

	err := database.UpdateCompany(companyName, newDetails)
	if err != nil {
		return c.Status(409).JSON(fiber.Map{
			"status": "error",
			"mesage": "resource conflict",
			"data":   err,
		})
	}

	return c.Status(204).JSON(fiber.Map{
		"status": "success",
		"mesage": "company updated",
		"data":   nil,
	})
}

func DeleteCompany(c *fiber.Ctx) error {
	// Handle Params
	companyName := c.Params("name")

	err := database.DeleteCompany(companyName)
	if err != nil {
		return c.Status(409).JSON(fiber.Map{
			"status":  "error",
			"message": "company not deleted",
			"data":    err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"mesage": "company deleted",
		"data":   nil,
	})
}
