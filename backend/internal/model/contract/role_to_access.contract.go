package contract

import (
	"context"
	"github.com/Halalins/backend/internal/model/entity"

	"github.com/google/uuid"
)

type RoleToAccessRepository interface {
	FindByRoleIDAndAccessID(ctx context.Context, roleID, accessID uuid.UUID) (*entity.RoleToAccess, error)
}
