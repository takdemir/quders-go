package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"errors"
	"gorm.io/gorm"
)

type QuestionAnswerRepository interface {
	FetchQuestionAnswers() ([]models.QuestionAnswer, error)
	FetchQuestionAnswerById(questionAnswerId uint) (*models.QuestionAnswer, error)
	CreateQuestionAnswer(newQuestionAnswer models.QuestionAnswer) error
	UpdateQuestionAnswer(questionAnswerId uint, newQuestionAnswer models.QuestionAnswer) error
}

type QuestionAnswerStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewQuestionAnswerStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *QuestionAnswerStore {
	return &QuestionAnswerStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (qa *QuestionAnswerStore) FetchQuestionAnswers() ([]models.QuestionAnswer, error) {
	var questionAnswers []models.QuestionAnswer
	result := qa.db.Preload("User").Preload("Question").Find(&questionAnswers)
	if result.Error != nil {
		return nil, result.Error
	}
	return questionAnswers, nil
}

func (qa *QuestionAnswerStore) FetchQuestionAnswerById(questionAnswerId uint) (*models.QuestionAnswer, error) {
	var questionAnswer models.QuestionAnswer
	result := qa.db.Preload("User").Preload("Question").Where("id=?", questionAnswerId).Find(&questionAnswer)
	if result.Error != nil {
		return nil, result.Error
	}
	if questionAnswer.ID == 0 {
		return nil, errors.New("question answer not found")
	}
	return &questionAnswer, nil
}

func (qa *QuestionAnswerStore) CreateQuestionAnswer(newQuestionAnswer models.QuestionAnswer) error {
	var questionAnswer models.QuestionAnswer
	result := qa.db.Where("user_id=? and question_id=?", newQuestionAnswer.UserId, newQuestionAnswer.QuestionId).Find(&questionAnswer)
	if result.Error != nil {
		return result.Error
	}
	if questionAnswer.ID > 0 {
		return errors.New("question answer for this user is already exist")
	}
	result = qa.db.Save(&newQuestionAnswer)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("could not create question answer")
	}
	return nil
}

func (qa *QuestionAnswerStore) UpdateQuestionAnswer(questionAnswerId uint, newQuestionAnswer models.QuestionAnswer) error {
	var questionAnswer models.QuestionAnswer
	result := qa.db.Where("id=?", questionAnswerId).Find(&questionAnswer)
	if result.Error != nil {
		return result.Error
	}
	if questionAnswer.ID == 0 {
		return errors.New("question answer not found")
	}

	result = qa.db.Model(&questionAnswer).Updates(newQuestionAnswer)
	if result.Error != nil {
		return result.Error
	}
	/*	if result.RowsAffected == 0 {
		return errors.New("no update for answer question")
	}*/
	return nil
}
