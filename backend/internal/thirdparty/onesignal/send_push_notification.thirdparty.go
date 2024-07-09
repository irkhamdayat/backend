package onesignal

import (
	"context"
	"fmt"
	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/request"

	"strings"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *ThirdParty) SendPushNotification(ctx context.Context, req request.EnqueueCreateNotificationReq,
	notificationID uuid.UUID) (err error) {
	OnesignalPushNotificationReq, _ := s.processTranslatePlaceHolder(req, notificationID)

	notification := s.createNotificationConfig(*OnesignalPushNotificationReq)
	appAuth := context.WithValue(ctx, onesignal.AppAuth, config.Env.Onesignal.ApiKey)

	resp, r, err := s.onesignalClient.DefaultApi.CreateNotification(appAuth).Notification(notification).Execute()
	if err != nil {
		logrus.Errorf("Error: %v\n", err)
		logrus.Errorf("Full HTTP response: %v\n", r)
		return err
	}

	logrus.Infof("Create Notification Resp: %v\n", resp)
	logrus.Info("Resp: ", util.Dump(resp))

	return
}

func (s *ThirdParty) createNotificationConfig(req request.OnesignalPushNotificationReq) onesignal.Notification {
	notification := *onesignal.NewNotification(config.Env.Onesignal.ApiID)

	notification.SetIncludeExternalUserIds(req.IncludeExternalUserIDs)
	notification.SetHeadings(req.Headings)
	notification.SetContents(req.Contents)
	notification.SetData(req.Data.MakeMap())

	// set rich media notification
	notification.SetBigPicture(req.Picture.AndroidURL)
	notification.SetIosAttachments(map[string]interface{}{
		"pic": req.Picture.IosURL,
	})
	notification.SetLargeIcon(req.Picture.LargeIcon)

	notification.SetIsIosNil()

	// make unique topic
	notification.SetWebPushTopic(uuid.New().String())

	return notification
}

func (s *ThirdParty) processTranslatePlaceHolder(req request.EnqueueCreateNotificationReq, notificationID uuid.UUID) (
	*request.OnesignalPushNotificationReq, error) {
	mapHeading := map[string]string{}
	mapContent := map[string]string{}

	for _, lang := range constant.Langs {
		lang = strings.ToLower(lang)
		text, err := util.TranslateWithLangAndPlaceholder(lang, s.i18nBundle, fmt.Sprintf("notifications.%s.Headline", req.ActionType), req.MessagePlaceHolder)
		if err != nil {
			return nil, err
		}
		mapHeading[lang] = text
		text, err = util.TranslateWithLangAndPlaceholder(lang, s.i18nBundle, fmt.Sprintf("notifications.%s.Content", req.ActionType), req.MessagePlaceHolder)
		if err != nil {
			return nil, err
		}
		mapContent[lang] = text
	}

	oneSignalHeading, err := s.parseOnesignalMap(mapHeading)
	if err != nil {
		return nil, err
	}

	oneSignalContent, err := s.parseOnesignalMap(mapContent)
	if err != nil {
		return nil, err
	}

	//WARNING: Change this external IDs
	var externalIDs = []string{
		"f2923a29-c806-4209-90dd-61f4f014ab03",
	}

	return &request.OnesignalPushNotificationReq{
		IncludeExternalUserIDs: externalIDs,
		Headings:               oneSignalHeading,
		Contents:               oneSignalContent,
		Data: request.AdditionalDataNotification{
			NotificationType: req.NotificationType,
			ActionType:       req.ActionType,
			NotificationID:   notificationID,
		},
	}, nil
}

func (s *ThirdParty) parseOnesignalMap(mapString map[string]string) (onesignal.StringMap, error) {
	res := onesignal.StringMap{}

	jsonData, err := json.Marshal(mapString)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
