package repository

import (
	"database/pkg/models"
	"gorm.io/gorm"
)

type CurrencyRepository interface {
	CreateCurrency(currency *models.Currency) *models.Currency
	GetCurrencies() []models.Currency
	GetCurrencyById(Id int64) (*models.Currency, *gorm.DB)
	GetCurrencyByCode(code string) (*models.Currency, *gorm.DB)
	DeleteCurrency(ID int64) *models.Currency
}

type CurrencyStore struct {
	db *gorm.DB
}

func NewCurrencyStore(db *gorm.DB) *CurrencyStore {
	return &CurrencyStore{
		db: db,
	}
}

func (currencyStore *CurrencyStore) CreateCurrency(currency *models.Currency) *models.Currency {
	currencyStore.db.Create(&currency)
	return currency
}

func (currencyStore *CurrencyStore) GetCurrencies() []models.Currency {
	var Currencies []models.Currency
	currencyStore.db.Find(&Currencies)
	return Currencies
}

func (currencyStore *CurrencyStore) GetCurrencyById(Id int64) (*models.Currency, *gorm.DB) {
	var currency *models.Currency
	db := currencyStore.db.Where("ID=?", Id).Find(&currency)
	return currency, db
}

func (currencyStore *CurrencyStore) GetCurrencyByCode(code string) (*models.Currency, *gorm.DB) {
	var currency *models.Currency
	db := currencyStore.db.Where("code=?", code).Find(&currency)
	return currency, db
}

func (currencyStore *CurrencyStore) DeleteCurrency(ID int64) *models.Currency {
	var currency *models.Currency
	result := currencyStore.db.Where("ID=?", ID).Find(&currency)
	if result.RowsAffected > 0 {
		currency.IsActive = false
		currencyStore.db.Save(currency)
	}
	return currency
}
