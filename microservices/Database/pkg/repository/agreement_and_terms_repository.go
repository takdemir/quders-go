package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"gorm.io/gorm"
)

type AgreementAndTermsRepository interface {
	FetchAgreementAndTerms() ([]models.AgreementAndTerms, error)
}

type AgreementAndTermsStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewAgreementAndTermsStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *AgreementAndTermsStore {
	return &AgreementAndTermsStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (atr *AgreementAndTermsStore) FetchAgreementAndTerms() ([]models.AgreementAndTerms, error) {
	var agreementsAndTerms []models.AgreementAndTerms
	result := atr.db.Find(&agreementsAndTerms)
	if result.Error != nil {
		return nil, result.Error
	}
	return agreementsAndTerms, nil
}
