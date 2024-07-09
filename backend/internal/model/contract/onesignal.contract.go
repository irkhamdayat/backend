package contract

import (
	"context"
	"github.com/google/uuid"

	"github.com/Halalins/backend/internal/model/request"
)

type OnesignalThirdParty interface {
	SendPushNotification(ctx context.Context, req request.EnqueuePushNotificationReq,
		notificationID uuid.UUID) (err error)
}
