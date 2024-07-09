package entity

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type Agent struct {
	ID                uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Code              string
	PhoneNumber       string
	BirthDate         time.Time
	BirthPlace        string
	Address           string
	Location          string
	Photo             uuid.NullUUID
	FirstName         string
	LastName          null.String
	Username          string
	Password          string
	Email             string
	Status            string
	CodeReferral      string
	KtpDocument       uuid.UUID
	KtpNumber         string
	NpwpDocument      uuid.UUID
	NpwpNumber        string
	BankAccountNumber string
	BankId            uuid.UUID
	Bank              Bank
	Pin               string
	IsSubscribeNews   bool
	ApprovedBy        uuid.NullUUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
