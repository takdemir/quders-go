package models

import "time"

type QuestionAnswer struct {
	ID         uint                 `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId     int                  `json:"userId"`
	User       User                 `json:"user" gorm:"foreignKey:UserId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	QuestionId int                  `json:"questionId"`
	Question   RegistrationQuestion `json:"question" gorm:"foreignKey:QuestionId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	IsChecked  *bool                `json:"isChecked" gorm:"not null"`
	CreatedAt  time.Time            `json:"createdAt" gorm:"autoCreateTime"`
}
