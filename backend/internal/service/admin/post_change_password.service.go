package admin

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) PostChangePassword(ctx context.Context, req request.AdminChangePasswordReq) (*response.IDResp, error) {
	var (
		tx     = s.db.WithContext(ctx).Begin()
		err    error
		logger = logrus.WithFields(logrus.Fields{
			"ctx":   util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"email": req.Email,
		})
	)

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logrus.Panic(p)
		}
		util.HandleTransaction(tx, err)
	}()

	admin, err := s.adminRepository.FindByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		logger.Errorf("failed to get admin: %v", err)
		return nil, err
	}

	token, err := util.GetCacheAndDelete[string](ctx, s.rdb, cachekey.AdminForgotPasswordTokenCacheKey(req.Email))
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Errorf("failed access redis: %v", err)
		return nil, err
	}

	if errors.Is(err, redis.Nil) || *token != req.Token {
		return nil, constant.ErrTokenNotMatchOrInvalid
	}

	admin.Password = util.EncryptWithSalt(req.Password, config.Env.Crypto.Salt)

	err = s.adminRepository.Update(ctx, admin)
	if err != nil {
		logger.Errorf("failed update admin: %v", err)
		return nil, err
	}

	return &response.IDResp{
		ID: admin.ID,
	}, nil
}
