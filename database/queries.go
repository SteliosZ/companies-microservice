package database

import (
	"company/microservice/model"

	"gorm.io/gorm"
)

// GetCompany queries postgres to get one company by name
func GetCompany(name string) *gorm.DB {
	db := DB.DB

	// Query DB
	var company = model.Company{Name: name}
	res := db.First(&company)

	return res
}
