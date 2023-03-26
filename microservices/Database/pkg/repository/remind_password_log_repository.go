package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"gorm.io/gorm"
)

type RemindPasswordLogRepository interface {
	CheckOneTimeCode(email string, oneTimeCode int) bool
	CreateOneTimeCode(newRemindPasswordLog models.RemindPasswordLog) error
	UpdateOneTimeCode(newRemindPasswordLog models.RemindPasswordLog) error
}

type RemindPasswordLogStore struct {
	db          *gorm.DB
	qudersRedis *redis.QudersRedis
}

func NewRemindPasswordLogStore(db *gorm.DB, qudersRedis *redis.QudersRedis) *RemindPasswordLogStore {
	return &RemindPasswordLogStore{
		db:          db,
		qudersRedis: qudersRedis,
	}
}

func (rpl *RemindPasswordLogStore) CheckOneTimeCode(email string, oneTimeCode int) bool {
	var user models.User
	var userCount int64
	rpl.db.Model(&user).Where("email=? and is_active=?", email, true).Find(&user).Count(&userCount)
	if userCount == 0 {
		return false
	}
	var remindPasswordLog models.RemindPasswordLog
	result := rpl.db.Where("email=? and one_time_code=? and is_used=?", email, oneTimeCode, false).Find(&remindPasswordLog)
	if result.Error != nil {
		return false
	}

	return true
}

func (rpl *RemindPasswordLogStore) CreateOneTimeCode(newRemindPasswordLog models.RemindPasswordLog) error {
	return nil
}

func (rpl *RemindPasswordLogStore) UpdateOneTimeCode(newRemindPasswordLog models.RemindPasswordLog) error {
	return nil
}
