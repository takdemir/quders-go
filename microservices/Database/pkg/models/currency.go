package models

import (
	_ "encoding/json"
	"time"
)

type Currency struct {
	//gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" gorm:"size:50;index;not null" validate:"required,min=2"`
	Code      string    `json:"code" gorm:"size:5;not null" validate:"required"`
	IsActive  *bool     `json:"isActive" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
