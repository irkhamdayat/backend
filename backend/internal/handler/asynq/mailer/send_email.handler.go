package mailer

import (
	"context"
	"errors"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/task"
)

func (h *Handler) SendEmail(ctx context.Context, t *asynq.Task) error {
	var (
		req    = request.SendEmailReq{}
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	err := util.BindingAsynqPayload(t, &req)
	if err != nil {
		logger.Errorf("failed binding async payload: %v", err)
		return err
	}

	switch t.Type() {
	case task.AsynqSendEmailBoilerplateTask:
		return h.mailerService.Send(ctx, req)
	default:
		return errors.New("invalid task type")
	}
}
