package entity

import "github.com/google/uuid"

type RoleToAccess struct {
	ID       uuid.UUID `gorm:"default:uuid_generate_v4()"`
	RoleID   uuid.UUID
	AccessID uuid.UUID
}
