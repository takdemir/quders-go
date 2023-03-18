package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type CurrencyRepository interface {
	CreateCurrency(currency models.Currency) error
	GetCurrencies() []models.Currency
	GetCurrencyById(Id int64) (*models.Currency, *gorm.DB)
	GetCurrencyByCode(code string) (*models.Currency, *gorm.DB)
	DeleteCurrency(ID int64) *models.Currency
}

type CurrencyStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewCurrencyStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *CurrencyStore {
	return &CurrencyStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (currencyStore *CurrencyStore) CreateCurrency(currency models.Currency) error {
	var existCurrency models.Currency
	var count int64
	err := currencyStore.db.Where("name=? or (code != '' and code != ?)", currency.Name, currency.Code).Find(&existCurrency).Count(&count)
	if err.Error != nil {
		return errors.New("currency check error" + err.Error.Error())
	}
	if count > 0 {
		return errors.New("currency is already exist")
	}
	result := currencyStore.db.Create(&currency)
	if result.Error != nil {
		return errors.New("create currency error: " + result.Error.Error())
	}
	if result.RowsAffected <= 0 {
		return errors.New("currency not created: " + result.Error.Error())
	}

	return nil
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
