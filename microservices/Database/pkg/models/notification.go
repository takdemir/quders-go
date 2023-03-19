package models

import "time"

type Notification struct {
	ID                  uint              `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	NotificationGroupId int               `json:"notificationGroupId"`
	NotificationGroup   NotificationGroup `json:"notificationGroup" gorm:"foreignKey:NotificationGroupId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	NotificationEventId int               `json:"notificationEventId"`
	NotificationEvent   NotificationEvent `json:"notificationEvent" gorm:"foreignKey:NotificationEventId;reference:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	MessageTemplate     int               `json:"messageTemplate"`
	Subject             string            `json:"subject" gorm:"type:text"`
	MessageText         string            `json:"messageText" gorm:"type:text"`
	Detail              string            `json:"detail" gorm:"type:text"`
	Type                string            `json:"type" gorm:"size:255"`
	Duration            string            `json:"duration" gorm:"size:50"`
	Severity            string            `json:"severity" gorm:"size:50"`
	RelatedTo           string            `json:"relatedTo" gorm:"size:255"`
	Region              string            `json:"region" gorm:"size:255"`
	IsNotified          bool              `json:"isNotified" gorm:"not null"`
	NotifiedDate        time.Time         `json:"notifiedDate"`
	ClosedAt            time.Time         `json:"closedAt"`
	CreatedAt           time.Time         `json:"createdAt" gorm:"autoCreateTime"`
}
