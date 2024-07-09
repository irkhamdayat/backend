package notification

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) BatchUpsertTranslate(ctx context.Context, translateNotificationHistories []entity.TranslateNotificationHistory) (err error) {
	if len(translateNotificationHistories) <= 0 {
		return nil
	}
	var (
		tx     = util.GetTxFromContext(ctx, r.db)
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
		})
	)

	err = tx.WithContext(ctx).
		Save(translateNotificationHistories).
		Error

	if err != nil {
		logger.Errorf("failed save batch translate: %v", err)
		return err
	}

	return nil
}
