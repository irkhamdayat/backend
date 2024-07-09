package notification

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) GetNotificationPagination(ctx context.Context, req request.GetNotificationPaginationReq) (
	result *response.GetNotificationPaginationResp, err error) {
	var (
		metaResp = req.Pagination.ToMetaResp()
		logger   = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
		})
		items = make([]response.NotificationPaginationResp, 0)
	)

	notifications, count, err := s.notificationHistoryRepository.GetPagination(ctx, req)
	if err != nil {
		logger.Errorf("failed get pagination: %v", err)
		return nil, err
	}

	for _, notification := range notifications {
		items = append(items, response.NotificationPaginationResp{
			NotificationType: notification.NotificationType,
			ActionType:       notification.ActionType,
			Headline:         notification.Headline,
			Message:          notification.Message,
			IsRead:           notification.IsRead,
			AdditionalData:   notification.AdditionalData,
			CreatedAt:        notification.CreatedAt.Format(time.RFC3339),
		})
	}

	return response.NewPaginationResp[response.NotificationPaginationResp](items, count, metaResp), nil
}
