package models

import "time"

type Framework struct {
	ID                   uint               `json:"id" gorm:"primaryKey;autoIncrement"`
	ProgramingLanguageId int                `json:"programingLanguageId"`
	ProgramingLanguage   ProgramingLanguage `json:"programingLanguage" gorm:"foreignKey:ProgramingLanguageId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Name                 string             `json:"name" gorm:"size:255;not null" validate:"required,min=2"`
	IsActive             *bool              `json:"isActive" gorm:"not null"`
	CreatedAt            time.Time          `json:"createdAt" gorm:"autoCreateTime"`
}
