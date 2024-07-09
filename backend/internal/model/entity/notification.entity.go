package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type NotificationHistory struct {
	ID               uuid.UUID `gorm:"default:uuid_generate_v4()"`
	IsRead           bool
	Image            string
	ActionType       string
	NotificationType string
	AdditionalData   json.RawMessage
}

type TranslateNotificationHistory struct {
	ID                    uuid.UUID `gorm:"default:uuid_generate_v4()"`
	NotificationHistoryID uuid.UUID
	Language              string
	Headline              string
	Message               string
}

type GetNotificationHistory struct {
	ID               uuid.UUID `gorm:"default:uuid_generate_v4()"`
	IsRead           bool
	Image            string
	ActionType       string
	NotificationType string
	AdditionalData   json.RawMessage
	Language         string
	Headline         string
	Message          string
	CreatedAt        time.Time
}
