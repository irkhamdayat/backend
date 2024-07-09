package entity

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Admin struct {
	ID               uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Photo            uuid.NullUUID
	FirstName        string
	LastName         null.String
	Username         string
	Password         string
	Email            string
	RoleID           uuid.UUID
	Role             Role
	InsuranceBrandID null.String
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
