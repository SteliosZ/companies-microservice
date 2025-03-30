package database

import (
	"company/microservice/model"
	"company/microservice/utils"
)

// GetCompany queries postgres to get one company by name
func GetCompany(companyName string) (error, model.Company) {
	var company = model.Company{}

	// Query DB
	res := tx.DB.First(&company, model.Company{Name: companyName})

	return res.Error, company
}

// CreateCompnay queries to create a new company entry
func CreateCompany(company model.Company) error {
	err := tx.DB.Create(&company).Error

	return err
}

// DeleteCompany soft deletes records from Company Table
func DeleteCompany(name string) error {
	var company = model.Company{Name: name}

	// Query DB
	res := tx.DB.Where(&company).Delete(&model.Company{})

	return res.Error
}

func UpdateCompany(companyName string, updatedFields map[string]interface{}) error {
	var company = model.Company{}
	res := tx.DB.First(&company, model.Company{Name: companyName})
	if res != nil {
		return res.Error
	}

	res = tx.DB.Model(&company).Updates(updatedFields)

	return res.Error
}

func CreateUser(user model.User) error {
	err := tx.DB.Create(&user).Error

	return err
}

func GetUserByEmail(u model.User) (model.User, error) {

	var user = model.User{}
	err := tx.DB.First(&user, model.User{Email: u.Email}).Error

	if err != nil {
		return model.User{}, err
	}

	err = utils.ComparePassword(user.Password, u.Password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
