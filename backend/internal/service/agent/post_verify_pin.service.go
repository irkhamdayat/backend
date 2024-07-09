package agent

import (
	"context"
	"errors"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (s *Service) PostVerifyPin(ctx context.Context, req request.PostVerifyPinReq) (*response.PostVerifyPinResp, error) {
	var (
		encryptedPin = util.EncryptWithSalt(req.Pin, config.Env.Crypto.Salt)
		logger       = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
		})
	)

	agent, err := s.agentRepository.FindByIDAndPin(ctx, req.ID, encryptedPin)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("error get agent: %v", err)
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = constant.ErrPinInvalid
		return nil, err
	}

	return &response.PostVerifyPinResp{
		ID: agent.ID,
	}, nil
}
