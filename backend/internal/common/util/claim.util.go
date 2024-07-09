package util

import (
	"context"

	"github.com/google/uuid"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/entity"
)

func GetAdminClaimFromContext(ctx context.Context, entType string) (*entity.AdminClaim, error) {
	id, ok := ctx.Value(constant.IDKey).(uuid.UUID)
	if !ok {
		return nil, constant.ErrUnauthorized
	}

	ent, ok := ctx.Value(constant.EntityKey).(string)
	if !ok {
		return nil, constant.ErrUnauthorized
	}

	if ent != entType {
		return nil, constant.ErrUnauthorized
	}

	roleID, ok := ctx.Value(constant.RoleIDKey).(uuid.UUID)
	if !ok {
		return nil, constant.ErrUnauthorized
	}

	brandInsuranceID, ok := ctx.Value(constant.BrandInsuranceKey).(uuid.NullUUID)
	if !ok {
		return nil, constant.ErrUnauthorized
	}

	return &entity.AdminClaim{
		ID:               id,
		RoleId:           roleID,
		InsuranceBrandId: brandInsuranceID,
	}, nil
}

func GetUserIDFromContext(ctx context.Context, entType string) (*uuid.UUID, error) {
	id, ok := ctx.Value(constant.IDKey).(uuid.UUID)
	if !ok {
		return nil, constant.ErrUnauthorized
	}

	ent, ok := ctx.Value(constant.EntityKey).(string)
	if !ok {
		return nil, constant.ErrUnauthorized
	}

	if ent != entType {
		return nil, constant.ErrUnauthorized
	}

	return &id, nil
}
