package admin

import (
	"context"
	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/sirupsen/logrus"
	"strings"
)

func (s *Service) PostRegister(ctx context.Context, req request.AdminRegisterReq) (resp *response.IDResp, err error) {
	var (
		tx     = s.db.WithContext(ctx).Begin()
		logger = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"username": req.Username,
			"email":    req.Email,
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

	encryptedPassword := util.EncryptWithSalt(req.Password, config.Env.Crypto.Salt)
	admin := &entity.Admin{
		Photo:            req.Photo,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Username:         strings.ToLower(req.Username),
		Password:         encryptedPassword,
		Email:            req.Email,
		RoleID:           req.RoleID,
		InsuranceBrandID: req.InsuranceBrandID,
		Status:           constant.AccountAdminStatusVerified,
	}

	err = s.adminRepository.Create(ctx, admin)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	//TODO: add send email to admin
	return &response.IDResp{
		ID: admin.ID,
	}, nil
}
