package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	FetchNotifications() (*[]models.Notification, error)
	FetchActiveNotifications() (*[]models.Notification, error)
	FetchNotificationById(id int) (*models.Notification, error)
}

type NotificationStore struct {
	db          *gorm.DB
	QudersRedis *redis.QudersRedis
}

func NewNotificationStore(db *gorm.DB, QudersRedis *redis.QudersRedis) *NotificationStore {
	return &NotificationStore{
		db:          db,
		QudersRedis: QudersRedis,
	}
}

func (NotificationStore *NotificationStore) FetchNotifications() (*[]models.Notification, error) {
	var Notification []models.Notification
	err := NotificationStore.db.Preload("NotificationGroup").Preload("NotificationEvent").Find(&Notification).Error
	if err != nil {
		return nil, err
	}
	return &Notification, nil
}

func (NotificationStore *NotificationStore) FetchActiveNotifications() (*[]models.Notification, error) {
	var Notification []models.Notification
	err := NotificationStore.db.Where("is_active=?", true).Preload("NotificationGroup").Preload("NotificationEvent").Find(&Notification).Error
	if err != nil {
		return nil, err
	}
	return &Notification, nil
}

func (NotificationStore *NotificationStore) FetchNotificationById(id int) (*models.Notification, error) {
	var Notification models.Notification
	err := NotificationStore.db.Where("id=?", id).Preload("NotificationGroup").Preload("NotificationEvent").Find(&Notification).Error
	if err != nil {
		return nil, err
	}
	return &Notification, nil
}
