package database

import (
	"company/microservice/model"
	"company/microservice/utils"

	"github.com/gofrs/uuid/v5"
)

// GetCompany queries postgres to get one company by name
func GetCompany(companyID *uuid.UUID) (error, *model.Company) {
	var company = model.Company{}

	// Query DB
	res := tx.DB.First(&company, model.Company{ID: *companyID})

	return res.Error, &company
}

// CreateCompnay queries to create a new company entry
func CreateCompany(company *model.Company) error {
	err := tx.DB.Create(&company).Error

	return err
}

// DeleteCompany soft deletes records from Company Table
func DeleteCompany(companyID *uuid.UUID) error {
	var company = model.Company{ID: *companyID}

	// Query DB
	res := tx.DB.Where(&company).Delete(&model.Company{})

	return res.Error
}

func UpdateCompany(companyID *uuid.UUID, updatedFields *model.Company) error {
	var company = model.Company{}

	res := tx.DB.First(&company, model.Company{ID: *companyID}).Error
	if res != nil {
		return res
	}

	res = tx.DB.Model(&company).Updates(&updatedFields).Error

	return res
}

func CreateUser(user *model.User) error {
	err := tx.DB.Create(&user).Error

	return err
}

func GetUserByEmail(u *model.User) (*model.User, error) {

	var user = model.User{}
	err := tx.DB.First(&user, model.User{Email: u.Email}).Error

	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(user.Password, u.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
