package admin

import (
	"context"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/Halalins/backend/internal/model/task"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) PostForgotPassword(ctx context.Context, req request.AdminForgotPasswordReq) (*response.EmailResp, error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":   util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"email": req.Email,
		})
	)

	admin, err := s.adminRepository.FindByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		logger.Errorf("failed to get admin: %v", err)
		return nil, err
	}

	token := util.GenerateRandomString(40, constant.AlphaNumeric)
	err = util.SetCache(ctx, s.rdb, cachekey.AdminForgotPasswordTokenCacheKey(req.Email), constant.TokenCacheDuration,
		token)
	if err != nil {
		logger.Errorf("failed setting cache redis: %v", err)
		return nil, err
	}

	err = util.ProcessPayloadAndEnqueueTask(s.asynqClient, task.AsynqSendEmailBoilerplateTask, request.SendEmailReq{
		Template: constant.MailerForgotPasswordTemplate,
		Subject:  constant.ForgotPasswordTemplate,
		To:       admin.Email,
		EmailBody: map[string]string{
			"FirstName": admin.FirstName,
			"LastName":  admin.LastName.String,
			"Token":     token,
		},
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"task": task.AsynqSendEmailBoilerplateTask,
		}).Error("failed send enqueue send email", err)
		return nil, err
	}

	return &response.EmailResp{
		Email: admin.Email,
	}, nil
}
