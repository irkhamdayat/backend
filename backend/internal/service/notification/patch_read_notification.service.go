package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) PatchReadNotification(ctx context.Context, req request.PatchReadNotificationReq) (
	resp response.IDResp, err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
		})
	)

	var notificationHistory *entity.NotificationHistory
	notificationHistory, err = s.notificationHistoryRepository.FindByID(ctx, uuid.MustParse(req.NotificationID))
	if err != nil {
		logger.Errorf("failed find by id notification history: %v", err)
		return
	}

	notificationHistory.IsRead = true

	err = s.notificationHistoryRepository.Upsert(ctx, notificationHistory)
	if err != nil {
		logger.Errorf("failed update notification history: %v", err)
		return
	}

	return response.IDResp{ID: notificationHistory.ID}, err
}
