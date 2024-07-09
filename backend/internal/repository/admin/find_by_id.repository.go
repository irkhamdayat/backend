package admin

import (
	"context"
	"github.com/Halalins/backend/internal/common/util"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (admin *entity.Admin, err error) {
	var (
		db     = util.GetTxFromContext(ctx, r.db)
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"ID":  ID.String(),
		})
	)

	err = db.WithContext(ctx).
		Where("id = ?", ID).
		Preload("Role").
		First(&admin).
		Error

	if err != nil {
		logger.Errorf("failed to find admin: %v", err)
		return
	}

	return
}
