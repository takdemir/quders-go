package models

import "time"

type NotificationGroup struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name          string    `json:"name" gorm:"size:255;index;not null"`
	Channel       string    `json:"channel" gorm:"size:50;not null"`
	ChannelValues []string  `json:"channelValues" gorm:"serializer:json"`
	Description   string    `json:"description" gorm:"type:text"`
	IsActive      bool      `json:"isActive" gorm:"not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

func (notificationGroup *NotificationGroup) Init() *NotificationGroup {
	notificationGroup.IsActive = true
	return notificationGroup
}
