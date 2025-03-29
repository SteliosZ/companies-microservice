package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Custom Type for Companies
type CompanyType string

// Corporations | NonProfit | Cooperative | Sole Proprietorship
const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

/*
Company Struct
@ID - uuid - required
@Name - 15 Chars - required - unique
@Description - 3000 Characters - optional
@Amount of Employees - int - required
@Registered - bool - required
@Type - CompanyType - required
*/
type Company struct {
	gorm.Model
	ID                uuid.UUID   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name              string      `gorm:"uniqueIndex;type:varchar(15);not null" json:"name"`
	Description       string      `gorm:"type:varchar(3000)" json:"description,omitempty"`
	AmountOfEmployees int         `gorm:"type:integer;not null" json:"amount_of_employees"`
	Registered        bool        `gorm:"type:boolean;not null" json:"registered"`
	Type              CompanyType `gorm:"type:varchar(50);not null" json:"type"`
}

func (user *Company) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}
