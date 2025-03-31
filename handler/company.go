package handler

import (
	"company/microservice/database"
	"company/microservice/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid/v5"
)

func GetCompanyByID(c *fiber.Ctx) error {
	// Handle Params
	companyUUID, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "company id is wrong",
			"data":    nil,
		})
	}

	err, res := database.GetCompany(companyUUID)
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
			"message": "company creation successful",
			"data":    err,
		})
}

func UpdateCompany(c *fiber.Ctx) error {
	var newDetails map[string]any

	companyUUID, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "company id is wrong",
			"data":    nil,
		})
	}

	// Parse and Check Body
	if err := c.BodyParser(&newDetails); err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"status":  "error",
				"message": "wrong input data",
				"data":    err,
			})
	}

	err = database.UpdateCompany(companyUUID, newDetails)
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
	companyUUID, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "company id is wrong",
			"data":    nil,
		})
	}

	err = database.DeleteCompany(companyUUID)
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
