package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type FrameworkRepository interface {
	FetchFrameworks() ([]models.Framework, error)
	FetchFrameworkById(FrameworkId uint) (*models.Framework, error)
	CreateFramework(newFramework models.Framework) error
	UpdateFramework(FrameworkId uint, newFramework models.Framework) error
}

type FrameworkStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewFrameworkStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *FrameworkStore {
	return &FrameworkStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (pl *FrameworkStore) FetchFrameworks() ([]models.Framework, error) {
	var Frameworks []models.Framework
	result := pl.db.Preload("ProgramingLanguage").Find(&Frameworks)
	if result.Error != nil {
		return nil, result.Error
	}
	return Frameworks, nil
}

func (pl *FrameworkStore) FetchFrameworkById(frameworkId uint) (*models.Framework, error) {
	var Framework models.Framework
	result := pl.db.Preload("ProgramingLanguage").Where("id=?", frameworkId).First(&Framework)
	if result.Error != nil {
		return nil, result.Error
	}
	if Framework.ID == 0 {
		return nil, errors.New("framework not found")
	}
	return &Framework, nil
}

func (pl *FrameworkStore) CreateFramework(newFramework models.Framework) error {
	var Framework models.Framework
	result := pl.db.Where("name=?", newFramework.Name).Find(&Framework)
	if result.Error != nil {
		return result.Error
	}
	if Framework.ID > 0 {
		return errors.New("framework is already exist")
	}
	result = pl.db.Save(&newFramework)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return result.Error
	}
	return nil
}

func (pl *FrameworkStore) UpdateFramework(FrameworkId uint, newFramework models.Framework) error {
	var Framework models.Framework
	result := pl.db.Where("name=?", newFramework.Name).Find(&Framework)
	if result.Error != nil {
		return result.Error
	}
	if Framework.ID > 0 && Framework.ID != FrameworkId {
		return errors.New("framework is already exist")
	}
	result = pl.db.Where("id=?", FrameworkId).Find(&Framework)
	if result.Error != nil {
		return result.Error
	}
	if Framework.ID == 0 {
		return errors.New("framework not found")
	}
	result = pl.db.Model(&Framework).Updates(&newFramework)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return result.Error
	}
	return nil
}
