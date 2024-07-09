package admin

import (
	"context"
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (s *Service) PostLogin(ctx context.Context, req request.AdminLoginReq) (*entity.UserClaim, error) {
	var (
		encryptedPassword = util.EncryptWithSalt(req.Password, config.Env.Crypto.Salt)
		logger            = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"username": req.Username,
		})
	)

	admin, err := s.adminRepository.FindByUsernameAndEncryptedPassword(ctx, strings.ToLower(req.Username), encryptedPassword)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrUnauthorized
	}

	if err != nil {
		logger.Errorf("failed get admin: %v", err)
		return nil, err
	}

	return &entity.UserClaim{
		ID:     admin.ID,
		Entity: constant.EntityTypeAdmin,
	}, nil
}
