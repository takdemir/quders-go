package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type CountryRepository interface {
	FetchCountries() ([]models.Country, error)
	FetchCountryById(countryId int) (*models.Country, error)
	CreateCountry(newCountry models.Country) error
	UpdateCountry(countryId int, newCountry models.Country) error
}

type CountryStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewCountryStore(
	db *gorm.DB,
	qudersRedis *redis.QudersRedis) *CountryStore {
	return &CountryStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (country *CountryStore) FetchCountries() ([]models.Country, error) {
	var countries []models.Country
	result := country.db.Preload("Currency").Find(&countries)
	if result.Error != nil {
		return nil, errors.New("fetch countries error: " + result.Error.Error())
	}
	return countries, nil
}

func (country *CountryStore) FetchCountryById(countryId int) (*models.Country, error) {
	var countryDetail models.Country
	result := country.db.Where("id=?", countryId).Preload("Currency").Find(&countryDetail)
	if result.Error != nil {
		return nil, errors.New("fetch countries error: " + result.Error.Error())
	}
	if countryDetail.ID == 0 {
		return nil, errors.New("no country found")
	}
	return &countryDetail, nil
}

func (country *CountryStore) CreateCountry(newCountry models.Country) error {
	result := country.db.Create(&newCountry)
	if result.Error != nil {
		return errors.New("create country error: " + result.Error.Error())
	}
	if result.RowsAffected <= 0 {
		return errors.New("create country error: " + result.Error.Error())
	}
	return nil
}

func (country *CountryStore) UpdateCountry(countryId int, newCountry models.Country) error {
	var countryDetail models.Country
	result := country.db.Where("id=?", countryId).First(&countryDetail)
	if result.Error != nil {
		return errors.New("update country error: " + result.Error.Error())
	}
	if countryDetail.ID == 0 {
		return errors.New("no country found")
	}
	countryDetail.Name = newCountry.Name
	countryDetail.Code = newCountry.Code
	countryDetail.IsActive = newCountry.IsActive
	result = country.db.Model(&countryDetail).Updates(countryDetail)
	if result.Error != nil {
		return errors.New("update country error: " + result.Error.Error())
	}

	return nil
}
