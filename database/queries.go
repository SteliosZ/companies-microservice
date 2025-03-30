package database

import (
	"company/microservice/model"
)

// GetCompany queries postgres to get one company by name
func GetCompany(name string) (error, model.Company) {
	db := DB.DB
	var company = model.Company{}

	// Query DB
	res := db.First(&company, model.Company{Name: name})

	return res.Error, company
}

// CreateCompnay queries to create a new company entry
func CreateCompany(company model.Company) error {
	db := DB.DB
	err := db.Create(&company).Error

	return err
}

// DeleteCompany soft deletes records from Company Table
func DeleteCompany(name string) error {
	db := DB.DB

	var company = model.Company{Name: name}

	// Query DB
	res := db.Where(&company).Delete(&model.Company{})

	return res.Error
}
