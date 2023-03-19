package models

import "time"

type AgreementAndTerms struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;"`
	Content   string    `json:"content" gorm:"type:text;not null" validate:"required,min=10"`
	IsActive  bool      `json:"isActive" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
