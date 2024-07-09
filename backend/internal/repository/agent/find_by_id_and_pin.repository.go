package agent

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByIDAndPin(ctx context.Context, id uuid.UUID, encryptedPin string) (
	agent *entity.Agent, err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"id":  id,
		})
		db = util.GetTxFromContext(ctx, r.db)
	)

	err = db.WithContext(ctx).
		Where("id = ?", id).
		Where("pin = ?", encryptedPin).
		First(&agent).
		Error

	if err != nil {
		logger.Errorf("failed to get agent: %v", err)
		return nil, err
	}

	return agent, nil
}
