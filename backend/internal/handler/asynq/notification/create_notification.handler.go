package notification

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (h *Handler) CreateNotification(ctx context.Context, t *asynq.Task) error {
	var (
		req    = request.EnqueueCreateNotificationReq{}
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	err := util.BindingAsynqPayload(t, &req)
	if err != nil {
		logger.Errorf("failed binding async payload: %v", err)
		return err
	}

	err = h.notificationService.CreateNotification(ctx, req)
	if err != nil {
		return err
	}

	return err
}
