package mailer

import (
	"context"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/sirupsen/logrus"
)

func (s *Service) Send(ctx context.Context, req request.SendEmailReq) (err error) {
	var (
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	err = s.mailerThirdParty.Send(ctx, req)
	if err != nil {
		logger.Errorf("failed send email: %v", err)
		return
	}

	return
}
