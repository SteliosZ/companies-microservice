package handler

import (
	"company/microservice/database"
	"company/microservice/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetCompanyByName(c *fiber.Ctx) error {
	// Handle Params
	companyName := c.Params("name")
	if companyName == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "company name cannot be empty",
			"data":    nil,
		})
	}

	err, res := database.GetCompany(companyName)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "could not get company",
			"data":    err,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"mesage": "company fetched",
		"data":   res,
	})
}

func CreateCompany(c *fiber.Ctx) error {
	company := model.Company{}

	// Parse and Check Body
	if err := c.BodyParser(&company); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong input data",
			"data":    err,
		})
	}

	// db Query for Creation
	err := database.CreateCompany(company)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"status":  "error",
				"message": "company not created",
				"data":    err,
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  "success",
			"message": "company create",
			"data":    err,
		})
}

func UpdateCompany(c *fiber.Ctx) error {
	newDetails := map[string]any{}

	companyName := c.Params("name")

	// Parse and Check Body
	if err := c.BodyParser(&newDetails); err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"status":  "error",
				"message": "wrong input data",
				"data":    err,
			})
	}

	err := database.UpdateCompany(companyName, newDetails)
	if err != nil {
		return c.Status(http.StatusConflict).
			JSON(fiber.Map{
				"status": "error",
				"mesage": "resource conflict",
				"data":   err,
			})
	}

	return c.Status(http.StatusNoContent).
		JSON(fiber.Map{
			"status": "success",
			"mesage": "company updated",
			"data":   nil,
		})
}

func DeleteCompany(c *fiber.Ctx) error {
	// Handle Params
	companyName := c.Params("name")
	if companyName == "" {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"status":  "error",
				"message": "company name cannot be empty",
				"data":    nil,
			})
	}

	err := database.DeleteCompany(companyName)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"status":  "error",
				"message": "company not deleted",
				"data":    err,
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status": "success",
			"mesage": "company deleted",
			"data":   nil,
		})
}
