package repository

import (
	"database/pkg/models"
	"database/pkg/utils/redis"
	"gorm.io/gorm"
)

type NotificationEventRepository interface {
	FetchNotificationEvents() (*[]models.NotificationEvent, error)
	FetchActiveNotificationEvents() (*[]models.NotificationEvent, error)
	FetchNotificationEventById(id int) (*models.NotificationEvent, error)
}

type NotificationEventStore struct {
	db          *gorm.DB
	QudersRedis *redis.QudersRedis
}

func NewNotificationEventStore(db *gorm.DB, QudersRedis *redis.QudersRedis) *NotificationEventStore {
	return &NotificationEventStore{
		db:          db,
		QudersRedis: QudersRedis,
	}
}

func (NotificationEventStore *NotificationEventStore) FetchNotificationEvents() (*[]models.NotificationEvent, error) {
	var NotificationEvent []models.NotificationEvent
	err := NotificationEventStore.db.Find(&NotificationEvent).Error
	if err != nil {
		return nil, err
	}
	return &NotificationEvent, nil
}

func (NotificationEventStore *NotificationEventStore) FetchActiveNotificationEvents() (*[]models.NotificationEvent, error) {
	var NotificationEvent []models.NotificationEvent
	err := NotificationEventStore.db.Where("is_active=?", true).Find(&NotificationEvent).Error
	if err != nil {
		return nil, err
	}
	return &NotificationEvent, nil
}

func (NotificationEventStore *NotificationEventStore) FetchNotificationEventById(id int) (*models.NotificationEvent, error) {
	var NotificationEvent models.NotificationEvent
	err := NotificationEventStore.db.Where("id=?", id).Find(&NotificationEvent).Error
	if err != nil {
		return nil, err
	}
	return &NotificationEvent, nil
}
