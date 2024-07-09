package response

import (
	"encoding/json"
)

type NotificationPaginationResp struct {
	NotificationType string          `json:"notificationType"`
	ActionType       string          `json:"actionType"`
	Headline         string          `json:"headline"`
	Message          string          `json:"message"`
	IsRead           bool            `json:"isRead"`
	AdditionalData   json.RawMessage `json:"additionalData"`
	CreatedAt        string          `json:"createdAt"`
}

type GetNotificationPaginationResp = PaginationResp[NotificationPaginationResp]
