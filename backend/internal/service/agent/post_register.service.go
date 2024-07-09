package agent

import (
	"context"
	"fmt"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/Halalins/backend/internal/model/task"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) PostRegister(ctx context.Context, req request.AgentRegisterReq) (resp *response.IDResp, err error) {
	var (
		tx     = s.db.WithContext(ctx).Begin()
		logger = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"username": req.Username,
			"email":    req.Email,
		})
		cacheKey = cachekey.AgentVerificationEmailTokenCacheKey(req.Email)
	)

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logrus.Panic(p)
		}
		util.HandleTransaction(tx, err)
	}()

	tokenCache, err := util.GetCacheAndDelete[string](ctx, s.rdb, cacheKey)

	if *tokenCache != req.OTP {
		err = constant.ErrOTPInvalid
		return nil, err
	}

	encryptedPassword := util.EncryptWithSalt(req.Password, config.Env.Crypto.Salt)
	encryptedPin := util.EncryptWithSalt(req.Pin, config.Env.Crypto.Salt)

	birthDate, err := time.Parse("02-01-2006", req.BirthDate)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	admin := &entity.Agent{
		Code:              fmt.Sprintf("HL-%d", time.Now().Unix()),
		PhoneNumber:       req.PhoneNumber,
		BirthDate:         birthDate,
		BirthPlace:        req.BirthPlace,
		Address:           req.Address,
		Location:          req.Location,
		Photo:             req.Photo,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		Username:          strings.ToLower(req.Username),
		Password:          encryptedPassword,
		Email:             req.Email,
		Status:            constant.AccountAdminStatusWaitingAdminApproved,
		CodeReferral:      util.GenerateRandomString(8, constant.AlphaCapitalNumeric),
		KtpDocument:       req.KtpDocument,
		KtpNumber:         req.KtpNumber,
		NpwpDocument:      req.NpwpDocument,
		NpwpNumber:        req.NpwpNumber,
		BankAccountNumber: req.BankAccountNumber,
		BankId:            req.BankId,
		Pin:               encryptedPin,
		IsSubscribeNews:   req.IsSubscribeNews,
	}

	err = s.agentRepository.Create(ctx, admin)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = util.ProcessPayloadAndEnqueueTask(s.asynqClient, task.AsynqSendEmailBoilerplateTask, request.SendEmailReq{
		Template: constant.MailerCreateAccountAgent,
		Subject:  constant.UploadSuccessSubject,
		To:       req.Email,
		EmailBody: map[string]string{
			"FirstName": req.FirstName,
			"LastName":  req.LastName.String,
		},
	})
	if err != nil {
		logger.Errorf("failed process queue email: %v", err)
		return nil, err
	}

	return &response.IDResp{
		ID: admin.ID,
	}, nil
}
