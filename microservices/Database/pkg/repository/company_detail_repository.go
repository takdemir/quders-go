package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type CompanyDetailRepository interface {
	FetchCompanyDetailByCompanyId(companyId uint) (models.CompanyDetail, error)
	CreateCompanyDetail(newCompanyDetail models.CompanyDetail) error
	UpdateCompanyDetail(id uint, newCompanyDetail models.CompanyDetail) error
}

type CompanyDetailStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewCompanyDetailStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *CompanyDetailStore {
	return &CompanyDetailStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (companyDetailStore *CompanyDetailStore) FetchCompanyDetailByCompanyId(companyId uint) (models.CompanyDetail, error) {
	var companyDetail models.CompanyDetail
	companyDetailStore.db.Where("company_id=?", companyId).Preload("Company").Preload("Country").Find(&companyDetail)
	return companyDetail, nil
}

func (companyDetailStore *CompanyDetailStore) CreateCompanyDetail(newCompanyDetail models.CompanyDetail) error {
	result := companyDetailStore.db.Save(&newCompanyDetail)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("not created company detail")
	}
	return nil
}

func (companyDetailStore *CompanyDetailStore) UpdateCompanyDetail(id uint, newCompanyDetail models.CompanyDetail) error {
	var companyDetail models.CompanyDetail
	result := companyDetailStore.db.Where("id=?", id).First(&companyDetail)
	if result.Error != nil {
		return result.Error
	}
	if companyDetail.ID == 0 {
		return errors.New("no company detail found")
	}
	companyDetailStore.db.Model(&companyDetail).Updates(newCompanyDetail)
	return nil
}
