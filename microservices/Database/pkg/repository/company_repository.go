package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	FetchCompanies() ([]models.Company, error)
	CreateCompany(newCompany models.Company) error
	UpdateCompany(companyId int, updatedCompanyDetail models.Company) error
	GetCompanyById(companyId int) (*models.Company, error)
}

type CompanyStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewCompanyStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *CompanyStore {
	return &CompanyStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (company *CompanyStore) FetchCompanies() ([]models.Company, error) {
	var companies []models.Company
	result := company.db.Find(&companies)
	if result.Error != nil {
		return nil, result.Error
	}
	return companies, nil
}

func (company *CompanyStore) CreateCompany(newCompany models.Company) error {
	result := company.db.Create(&newCompany)
	if result.Error != nil {
		return errors.New("error on company create: " + result.Error.Error())
	}

	if result.RowsAffected < 0 {
		return errors.New("no company created" + result.Error.Error())
	}
	return nil
}

func (company *CompanyStore) UpdateCompany(companyId int, updatedCompanyDetail models.Company) error {
	var companyDetail models.Company
	result := company.db.Where("id=?", companyId).Find(&companyDetail)

	if result.Error != nil {
		return errors.New("error on company update: " + result.Error.Error())
	}

	if companyDetail.ID == 0 {
		return errors.New("no company found")
	}
	companyDetail.Name = updatedCompanyDetail.Name
	companyDetail.IsActive = updatedCompanyDetail.IsActive
	result = company.db.Save(&companyDetail)
	if result.RowsAffected < 0 {
		return errors.New("company update error")
	}
	return nil
}

func (company *CompanyStore) GetCompanyById(companyId int) (*models.Company, error) {
	var companyDetail models.Company
	result := company.db.Where("id=?", companyId).Find(&companyDetail)

	if result.Error != nil {
		return nil, errors.New("error on company fetch: " + result.Error.Error())
	}

	if companyDetail.ID == 0 {
		return nil, errors.New("no company found")
	}
	return &companyDetail, nil
}
