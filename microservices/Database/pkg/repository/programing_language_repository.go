package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type ProgramingLanguageRepository interface {
	FetchProgramingLanguages() ([]models.ProgramingLanguage, error)
	FetchProgramingLanguageById(programingLanguageId uint) (*models.ProgramingLanguage, error)
	CreateProgramingLanguage(newProgramingLanguage models.ProgramingLanguage) error
	UpdateProgramingLanguage(programingLanguageId uint, newProgramingLanguage models.ProgramingLanguage) error
}

type ProgramingLanguageStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewProgramingLanguageStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *ProgramingLanguageStore {
	return &ProgramingLanguageStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (pl *ProgramingLanguageStore) FetchProgramingLanguages() ([]models.ProgramingLanguage, error) {
	var programingLanguages []models.ProgramingLanguage
	result := pl.db.Find(&programingLanguages)
	if result.Error != nil {
		return nil, result.Error
	}
	return programingLanguages, nil
}

func (pl *ProgramingLanguageStore) FetchProgramingLanguageById(programingLanguageId uint) (*models.ProgramingLanguage, error) {
	var programingLanguage models.ProgramingLanguage
	result := pl.db.Where("id=?", programingLanguageId).First(&programingLanguage)
	if result.Error != nil {
		return nil, result.Error
	}
	if programingLanguage.ID == 0 {
		return nil, errors.New("programing language not found")
	}
	return &programingLanguage, nil
}

func (pl *ProgramingLanguageStore) CreateProgramingLanguage(newProgramingLanguage models.ProgramingLanguage) error {
	var programingLanguage models.ProgramingLanguage
	result := pl.db.Where("name=?", newProgramingLanguage.Name).Find(&programingLanguage)
	if result.Error != nil {
		return result.Error
	}
	if programingLanguage.ID > 0 {
		return errors.New("programing language is already exist")
	}
	result = pl.db.Save(&newProgramingLanguage)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return result.Error
	}
	return nil
}

func (pl *ProgramingLanguageStore) UpdateProgramingLanguage(programingLanguageId uint, newProgramingLanguage models.ProgramingLanguage) error {
	var programingLanguage models.ProgramingLanguage
	result := pl.db.Where("name=?", newProgramingLanguage.Name).Find(&programingLanguage)
	if result.Error != nil {
		return result.Error
	}
	if programingLanguage.ID > 0 && programingLanguage.ID != programingLanguageId {
		return errors.New("programing language is already exist")
	}
	result = pl.db.Where("id=?", programingLanguageId).Find(&programingLanguage)
	if result.Error != nil {
		return result.Error
	}
	if programingLanguage.ID == 0 {
		return errors.New("programing language not found")
	}
	result = pl.db.Model(&programingLanguage).Updates(&newProgramingLanguage)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return result.Error
	}
	return nil
}
