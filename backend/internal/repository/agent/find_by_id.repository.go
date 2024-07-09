package agent

import (
	"context"
	"github.com/Halalins/backend/internal/common/util"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (agent *entity.Agent, err error) {
	var (
		db     = util.GetTxFromContext(ctx, r.db)
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"ID":  ID.String(),
		})
	)

	err = db.WithContext(ctx).
		Where("id = ?", ID).
		Preload("Bank").
		First(&agent).
		Error

	if err != nil {
		logger.Errorf("failed to find agent: %v", err)
		return
	}

	return
}
