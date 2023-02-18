package models

import (
	_ "encoding/json"
	"gorm.io/gorm"
	"quders/pkg/config"
)

var db *gorm.DB

type Currency struct {
	gorm.Model
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  string `json:"isActive"`
	CreatedAt string `json:"createdAt"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Currency{})
	if err != nil {
		return
	}
}

func (b *Currency) CreateCurrency() *Currency {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetCurrencies() []Currency {
	var Currencies []Currency
	db.Find(&Currencies)
	return Currencies
}

func GetCurrencyById(Id int64) (*Currency, *gorm.DB) {
	var currency Currency
	db := db.Where("ID=?", Id).Find(&currency)
	return &currency, db
}

func DeleteCurrency(ID int64) Currency {
	var currency Currency
	db.Where("ID=?", ID).Delete(currency)
	return currency
}
