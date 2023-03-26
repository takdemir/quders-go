package models

import "time"

type CompanyDetail struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CompanyId   int       `json:"companyId"`
	Company     Company   `json:"company" gorm:"foreignKey:CompanyId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CountryId   int       `json:"countryId"`
	Country     Country   `json:"country" gorm:"foreignKey:CountryId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Email       string    `json:"email" gorm:"size:255;not null" validate:"required,email"`
	Phone       string    `json:"phone" gorm:"size:255"`
	MobilePhone string    `json:"mobilePhone" gorm:"size:255"`
	Address     string    `json:"address" gorm:"type:text"`
	Address2    string    `json:"address2" gorm:"type:text"`
	IsActive    *bool     `json:"isActive" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
