package agent

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/Halalins/backend/internal/model/task"
)

func (s *Service) PostVerifyEmail(ctx context.Context, req request.PostVerifyEmailReq) (
	resp *response.PostVerifyEmailResp, err error) {
	var (
		tx     = s.db.WithContext(ctx).Begin()
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"req": util.Dump(req),
		})
		token    = util.GenerateRandomString(6, constant.Numeric)
		cacheKey = cachekey.AgentVerificationEmailTokenCacheKey(req.Email)
	)

	ctx = util.NewTxContext(ctx, tx)

	err = util.SetCache(ctx, s.rdb, cacheKey, constant.TokenCacheDuration,
		token)
	if err != nil {
		logger.Errorf("failed setting cache redis: %v", err)
		return
	}

	agent, err := s.agentRepository.FindByEmailOrUsername(ctx, req.Email, "")
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("failed find agent with email")
		return
	}

	if agent.ID != uuid.Nil {
		err = constant.ErrAccountAlreadyExist
		return
	}

	err = util.ProcessPayloadAndEnqueueTask(s.asynqClient, task.AsynqSendEmailBoilerplateTask, request.SendEmailReq{
		Template: constant.MailerVerifyEmailAgent,
		Subject:  constant.ForgotVerifyEmailAgent,
		To:       req.Email,
		EmailBody: map[string]string{
			"Token": token,
		},
	})
	if err != nil {
		logger.Errorf("failed process queue email: %v", err)
		return nil, err
	}

	resp = &response.PostVerifyEmailResp{Email: req.Email}
	return
}
