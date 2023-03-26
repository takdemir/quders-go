package models

import "time"

type Company struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" gorm:"type:text;not null" validate:"required,min=3"`
	IsActive  *bool     `json:"isActive" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
