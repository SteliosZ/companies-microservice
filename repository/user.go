package repository

import (
	"company/microservice/model"
	"company/microservice/utils"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(u *model.User) (*model.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{DB: db}
}

// CreateUser is used for registration and adds a new user entry
func (tx UserRepository) CreateUser(user *model.User) error {
	err := tx.DB.Create(&user).Error

	return err
}

// GetUserByEmail is used for login, retrieves the user from database
// and returns a jwt token
func (tx UserRepository) GetUserByEmail(u *model.User) (*model.User, error) {

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
