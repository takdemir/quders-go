package models

import "time"

type NotificationEvent struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" gorm:"size:255;index;not null"`
	IsActive  bool      `json:"isActive" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

func (notificationEvent *NotificationEvent) Init() *NotificationEvent {
	notificationEvent.IsActive = true
	return notificationEvent
}
