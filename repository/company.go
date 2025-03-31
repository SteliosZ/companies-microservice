package repository

import (
	"company/microservice/model"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type ICompanyRepository interface {
	GetByID(id *uuid.UUID) (error, *model.Company)
	Create(company *model.Company) error
	Delete(companyID *uuid.UUID) error
	Update(companyID *uuid.UUID, updatedFields *model.Company) error
}

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) ICompanyRepository {
	return &CompanyRepository{DB: db}
}

// GetCompany queries postgres to get one company by name
func (tx CompanyRepository) GetByID(companyID *uuid.UUID) (error, *model.Company) {
	var company = model.Company{}

	// Query DB
	res := tx.DB.First(&company, model.Company{ID: *companyID})

	return res.Error, &company
}

// CreateCompnay queries to create a new company entry
func (tx CompanyRepository) Create(company *model.Company) error {
	err := tx.DB.Create(&company).Error

	return err
}

// DeleteCompany soft deletes records from Company Table
func (tx CompanyRepository) Delete(companyID *uuid.UUID) error {
	var company = model.Company{ID: *companyID}

	// Query DB
	res := tx.DB.Where(&company).Delete(&model.Company{})

	return res.Error
}

// UpdateCompany updates an entry in database in any given field from model.Company
func (tx CompanyRepository) Update(companyID *uuid.UUID, updatedFields *model.Company) error {
	var company = model.Company{}

	res := tx.DB.First(&company, model.Company{ID: *companyID}).Error
	if res != nil {
		return res
	}

	res = tx.DB.Model(&company).Updates(&updatedFields).Error

	return res
}
