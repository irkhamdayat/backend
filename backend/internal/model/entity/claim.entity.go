package entity

import (
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/google/uuid"
)

type UserClaim struct {
	ID     uuid.UUID
	Entity string
}

type AdminClaim struct {
	ID               uuid.UUID
	RoleId           uuid.UUID
	InsuranceBrandId uuid.NullUUID
}

func ValidatorEntityAdmin(claim UserClaim) error {
	if claim.ID == uuid.Nil {
		return constant.ErrUnauthorized
	}

	if claim.Entity != constant.EntityTypeAdmin {
		return constant.ErrUnauthorized
	}

	return nil
}

func ValidatorEntityAgent(claim UserClaim) error {
	if claim.ID == uuid.Nil {
		return constant.ErrUnauthorized
	}

	if claim.Entity != constant.EntityTypeAgent {
		return constant.ErrUnauthorized
	}

	return nil
}
