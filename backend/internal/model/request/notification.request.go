package request

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"

	"github.com/OneSignal/onesignal-go-api"
)

type EnqueueCreateNotificationReq struct {
	NotificationType   string         `json:"notificationType"`
	ActionType         string         `json:"actionType"`
	Image              string         `json:"image"`
	MessagePlaceHolder map[string]any `json:"messagePlaceHolder"`
}

type EnqueuePushNotificationReq struct {
	MessagePlaceholder     map[string]any
	IncludeExternalUserIDs []string
	CompanyCrewID          uuid.UUID
	ActionType             string
	MessageID              string
	NotificationType       string
	MessagePlaceHolder     map[string]any
}

type OnesignalPushNotificationReq struct {
	IncludeExternalUserIDs []string
	Headings               onesignal.StringMap
	Contents               onesignal.StringMap
	Data                   AdditionalDataNotification
	Picture                OnesignalImage
}

type OnesignalImage struct {
	AndroidURL string
	IosURL     string
	LargeIcon  string
}

type AdditionalDataNotification struct {
	ActionType       string    `json:"actionType"`
	NotificationType string    `json:"notificationType"`
	NotificationID   uuid.UUID `json:"notificationID"`
}

func (a AdditionalDataNotification) toMap() map[string]interface{} {
	data := make(map[string]interface{})
	if a.ActionType != "" {
		data["actionType"] = a.ActionType
	}
	if a.NotificationType != "" {
		data["notificationType"] = a.NotificationType
	}

	return data
}

func (a AdditionalDataNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.toMap())
}

func (a AdditionalDataNotification) MakeMap() map[string]interface{} {
	return a.toMap()
}

type GetNotificationPaginationReq struct {
	CompanyID        uuid.UUID `json:"-"`
	CompanyUserID    uuid.UUID `json:"-"`
	CompanyCrewID    uuid.UUID `json:"-"`
	ContractPeriodID uuid.UUID `json:"-"`

	Pagination PaginationReq
}

type PatchReadNotificationReq struct {
	CompanyID        uuid.UUID `json:"-"`
	CompanyUserID    uuid.UUID `json:"-"`
	CompanyCrewID    uuid.UUID `json:"-"`
	ContractPeriodID uuid.UUID `json:"-"`
	NotificationID   string    `uri:"notificationID" binding:"required,uuid"`
}
