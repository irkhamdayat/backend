package roletoacess

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByRoleIDAndAccessID(ctx context.Context, roleID, accessID uuid.UUID) (*entity.RoleToAccess, error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"roleID":   roleID.String(),
			"accessID": accessID.String(),
		})
		db = util.GetTxFromContext(ctx, r.db)
	)
	roleToAccess := new(entity.RoleToAccess)

	err := db.WithContext(ctx).
		Where("role_id = ? AND access_id = ?", roleID, accessID).
		First(roleToAccess).
		Error

	if err != nil {
		logger.Errorf("failed get by role id and access id")
		return nil, err
	}

	return roleToAccess, nil
}
