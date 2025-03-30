package model

import (
	"company/microservice/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email    string    `gorm:"uniqueIndex;type:varchar(50);not null" json:"email"`
	Password string    `gorm:"type:varchar(128);not null" json:"password"`
}

// BeforeCreate handles email string check, password hashing and uuid creation
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid creation
	u.ID = uuid.New()

	// check email string validity
	if err := utils.CheckEmailValidity(u.Email); err != nil {
		return err
	}

	// hash password
	newHash, err := utils.GenerateHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = newHash

	return nil
}
