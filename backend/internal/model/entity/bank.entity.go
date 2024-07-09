package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bank struct {
	ID   uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Code string
	Icon uuid.UUID
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
