package models

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey,autoIncrement,not null"`
	FirstName   string    `json:"firstName" gorm:"size:50;not null"`
	LastName    string    `json:"lastName" gorm:"size:50;not null"`
	Email       string    `json:"email" gorm:"size:255;not null;unique"`
	Password    string    `json:"password" gorm:"size:255;not null;"`
	MobilePhone string    `json:"mobilePhone" gorm:"size:20;not null;"`
	IsActive    *bool     `json:"isActive" gorm:"not null;"`
	Roles       []string  `json:"roles" gorm:"serializer:json;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
