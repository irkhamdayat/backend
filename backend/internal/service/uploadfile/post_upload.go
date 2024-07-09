package uploadfile

import (
	"context"
	"github.com/Halalins/backend/internal/model/task"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) Upload(ctx context.Context, req request.UploadFileReq) (*response.IDResp, error) {
	var (
		tx               = s.db.Begin().WithContext(ctx)
		logger           = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
		err              error
		allowedFileTypes []string
	)

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logrus.Panic(p)
		}
		util.HandleTransaction(tx, err)
	}()

	switch req.UploadType {
	case constant.UploadTypeEvidance, constant.UploadTypeProfilePicture:
		allowedFileTypes = constant.FileTypeImage
	case constant.UploadTypeRichMedia:
		allowedFileTypes = constant.FileTypeRichMedia
	}

	fileName, err := s.cloudStorageThirdParty.Upload(ctx, req.File, allowedFileTypes, strings.ToLower(req.UploadType))
	if err != nil {
		logger.Errorf("failed upload to cloud storage: %v", err)
		return nil, err
	}

	fileID, err := s.uploadFileRepository.Create(ctx, entity.UploadFile{
		FileName:   fileName,
		UploadType: req.UploadType,
	})
	if err != nil {
		logger.Errorf("failed create upload file history: %v", err)
		return nil, err
	}

	err = util.ProcessPayloadAndEnqueueTask(s.asynqClient, task.AsynqSendEmailBoilerplateTask, request.SendEmailReq{
		Template: constant.MailerUploadSuccessTemplate,
		Subject:  constant.UploadSuccessSubject,
		To:       "avtara.id@gmail.com",
		EmailBody: map[string]string{
			"FileName":   fileName,
			"UploadDate": time.Now().Format(time.RFC850),
		},
	})
	if err != nil {
		logger.Errorf("failed process queue email: %v", err)
		return nil, err
	}

	err = util.ProcessPayloadAndEnqueueTask(s.asynqClient, task.AsynqCreateNotification,
		request.EnqueueCreateNotificationReq{
			NotificationType: constant.NotificationTypeNotify,
			ActionType:       constant.NotificationActionTypeSuccessUploadFile,
			MessagePlaceHolder: map[string]any{
				"brand": "Onesignal",
			},
		})
	if err != nil {
		logger.Errorf("failed process queue notification: %v", err)
		return nil, err
	}

	return &response.IDResp{ID: *fileID}, nil
}
