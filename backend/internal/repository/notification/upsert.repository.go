package notification

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) Upsert(ctx context.Context, notificationHistory *entity.NotificationHistory) (err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
		})
		tx = util.GetTxFromContext(ctx, r.db)
	)

	err = tx.WithContext(ctx).
		Save(notificationHistory).
		Error

	if err != nil {
		logger.Errorf("failed to upsert notification history: %v", err)
		return err
	}

	return nil
}
