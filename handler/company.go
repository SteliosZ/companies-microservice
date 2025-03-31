package handler

import (
	"company/microservice/model"
	"company/microservice/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid/v5"
)

type ICompanyHandler interface {
	GetCompanyByID(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
	UpdateCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
}

type CompanyHandler struct {
	CompanyRepository repository.ICompanyRepository
}

func NewCompanyHandler(repo repository.ICompanyRepository) ICompanyHandler {
	return &CompanyHandler{CompanyRepository: repo}
}

// GetCompanyByID retrieves a company entry from db
func (h CompanyHandler) GetCompanyByID(c *fiber.Ctx) error {
	// Handle Params
	companyUUID, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "company id is wrong",
			"data":    nil,
		})
	}

	err, res := h.CompanyRepository.GetByID(&companyUUID)
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

// CreateCompany adds a new entry of company to the db
func (h CompanyHandler) CreateCompany(c *fiber.Ctx) error {
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
	err := h.CompanyRepository.Create(&company)
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

// UpdateCompany handler updates a company entry with given data
func (h CompanyHandler) UpdateCompany(c *fiber.Ctx) error {
	newDetails := model.Company{}

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

	err = h.CompanyRepository.Update(&companyUUID, &newDetails)
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

// DeleteCompany handler contains the functionality of deleting a company entry from db
func (h CompanyHandler) DeleteCompany(c *fiber.Ctx) error {
	// Handle Params
	companyUUID, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "company id is wrong",
			"data":    nil,
		})
	}

	err = h.CompanyRepository.Delete(&companyUUID)
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
