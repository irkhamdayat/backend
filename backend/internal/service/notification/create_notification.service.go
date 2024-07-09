package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (s *Service) CreateNotification(ctx context.Context, req request.EnqueueCreateNotificationReq) (err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":                          util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"enqueueCreateNotificationReq": util.Dump(req),
		})
		tx = s.db.WithContext(ctx).Begin()
	)

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logrus.Panic(p)
		}
		util.HandleTransaction(tx, err)
	}()

	notification, err := s.createNotificationHistory(ctx, req)
	if err != nil {
		logger.Errorf("failed to create notification history: %v", err)
		return err
	}

	if err = s.notificationHistoryRepository.Upsert(ctx, notification); err != nil {
		logger.Errorf("failed to upsert notification history: %v", err)
		return err
	}

	if err = s.createTranslateNotifications(ctx, req, notification.ID); err != nil {
		logger.Errorf("failed to create translation notification: %v", err)
		return err
	}

	if err = s.onesignalThirdParty.SendPushNotification(ctx, req, notification.ID); err != nil {
		logger.Errorf("failed to send push notification: %v", err)
		return err
	}
	return
}

func (s *Service) createNotificationHistory(ctx context.Context, req request.EnqueueCreateNotificationReq) (notification *entity.NotificationHistory,
	err error) {
	additionalData, err := request.AdditionalDataNotification{}.MarshalJSON()
	if err != nil {
		return
	}

	notification = &entity.NotificationHistory{
		Image:            req.Image,
		ActionType:       req.ActionType,
		NotificationType: req.NotificationType,
		AdditionalData:   additionalData,
	}

	return
}

func (s *Service) createTranslateNotifications(ctx context.Context, req request.EnqueueCreateNotificationReq, notificationID uuid.UUID) error {
	var newNotifications []entity.TranslateNotificationHistory

	for _, lang := range constant.Langs {
		headline, message, err := s.translateNotification(lang, req)
		if err != nil {
			return err
		}

		newNotifications = append(newNotifications, entity.TranslateNotificationHistory{
			NotificationHistoryID: notificationID,
			Language:              lang,
			Headline:              headline,
			Message:               message,
		})
	}

	return s.notificationHistoryRepository.BatchUpsertTranslate(ctx, newNotifications)
}

func (s *Service) translateNotification(lang string, req request.EnqueueCreateNotificationReq) (headline, message string, err error) {
	headline, err = util.TranslateWithLangAndPlaceholder(lang, s.i18nBundle, fmt.Sprintf("notifications.%s.Headline", req.ActionType), req.MessagePlaceHolder)
	if err != nil {
		return
	}

	message, err = util.TranslateWithLangAndPlaceholder(lang, s.i18nBundle, fmt.Sprintf("notifications.%s.Message", req.ActionType), req.MessagePlaceHolder)
	if err != nil {
		return
	}

	return headline, message, nil
}
