package models

import "time"

type RemindPasswordLog struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int       `json:"userId"`
	User         User      `json:"user" gorm:"foreignKey:UserId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Email        string    `json:"email" gorm:"size:255;not null" validate:"required,email"`
	OneTimeCode  int       `json:"oneTimeCode" gorm:"not null" validate:"required"`
	IsUsed       *bool     `json:"isUsed" gorm:"not null"`
	CodeUsedDate time.Time `json:"codeUsedDate"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
