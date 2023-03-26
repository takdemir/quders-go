package models

import "time"

type ProgramingLanguage struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:255;not null" validate:"required,min=2"`
	Icon      string    `json:"icon" gorm:"size:255"`
	IsActive  *bool     `json:"isActive" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
