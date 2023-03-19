package models

type Country struct {
	ID         uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string   `json:"name" gorm:"size:255;not null" validate:"required,min=2"`
	Code       string   `json:"code" gorm:"size:10"`
	IsActive   bool     `json:"isActive" gorm:"not null"`
	CurrencyId int      `json:"currencyId"`
	Currency   Currency `json:"currency" gorm:"foreignKey:CurrencyId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}
