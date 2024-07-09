package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (*entity.NotificationHistory, error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"id":  ID.String(),
		})
		notification = new(entity.NotificationHistory)
	)

	err := r.db.WithContext(ctx).
		Where("id = ?", ID).
		First(notification).
		Error

	if err != nil {
		logger.Errorf("failed to find by id: %v", err)
		return nil, err
	}

	return notification, nil
}
