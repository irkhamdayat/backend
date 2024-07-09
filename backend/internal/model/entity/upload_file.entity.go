package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UploadFile struct {
	ID         uuid.UUID `gorm:"default:uuid_generate_v4()"`
	FileName   string
	UploadType string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
