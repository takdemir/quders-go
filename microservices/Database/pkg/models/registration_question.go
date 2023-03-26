package models

import "time"

const (
	EMPLOYEE string = "EMPLOYEE"
	EMPLOYER string = "EMPLOYER"
)

type RegistrationQuestion struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Question  string    `json:"question" gorm:"type:text;not null" validate:"required,min=10"`
	IsActive  *bool     `json:"isActive" gorm:"not null"`
	UserType  string    `json:"userType" gorm:"size:25;required;not null" validate:"oneof=EMPLOYEE EMPLOYER"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
