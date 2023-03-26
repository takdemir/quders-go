package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type RegistrationQuestionRepository interface {
	FetchRegistrationQuestions() ([]models.RegistrationQuestion, error)
	FetchRegistrationQuestionById(id uint) (*models.RegistrationQuestion, error)
	CreateRegistrationQuestion(newRegistrationQuestion models.RegistrationQuestion) error
	UpdateRegistrationQuestion(regId uint, newRegistrationQuestion models.RegistrationQuestion) error
}

type RegistrationQuestionStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewRegistrationQuestionStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *RegistrationQuestionStore {
	return &RegistrationQuestionStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (rq *RegistrationQuestionStore) FetchRegistrationQuestions() ([]models.RegistrationQuestion, error) {
	var registrationQuestions []models.RegistrationQuestion
	result := rq.db.Find(&registrationQuestions)
	if result.Error != nil {
		return nil, result.Error
	}
	return registrationQuestions, nil
}

func (rq *RegistrationQuestionStore) FetchRegistrationQuestionById(id uint) (*models.RegistrationQuestion, error) {
	var registrationQuestion models.RegistrationQuestion
	result := rq.db.Where("id=?", id).Find(&registrationQuestion)
	if result.Error != nil {
		return nil, result.Error
	}
	if registrationQuestion.ID == 0 {
		return nil, errors.New("registration question not found")
	}
	return &registrationQuestion, nil
}

func (rq *RegistrationQuestionStore) CreateRegistrationQuestion(newRegistrationQuestion models.RegistrationQuestion) error {
	result := rq.db.Save(&newRegistrationQuestion)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("could not create registration question")
	}
	return nil
}

func (rq *RegistrationQuestionStore) UpdateRegistrationQuestion(regId uint, newRegistrationQuestion models.RegistrationQuestion) error {
	var registrationQuestion models.RegistrationQuestion
	result := rq.db.Where("id=?", regId).Find(&registrationQuestion)
	if result.Error != nil {
		return result.Error
	}
	if registrationQuestion.ID == 0 {
		return errors.New("registration question not found")
	}
	result = rq.db.Model(&registrationQuestion).Updates(&newRegistrationQuestion)
	if result.Error != nil {
		return result.Error
	}
	/*	if result.RowsAffected == 0 {
		return errors.New("could not update registration question")
	}*/
	return nil
}
