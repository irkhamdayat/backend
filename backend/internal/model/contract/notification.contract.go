package contract

import (
	"context"

	"github.com/google/uuid"

	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, req request.EnqueueCreateNotificationReq) error
	GetNotificationPagination(ctx context.Context, req request.GetNotificationPaginationReq) (
		result *response.GetNotificationPaginationResp, err error)
	PatchReadNotification(ctx context.Context, req request.PatchReadNotificationReq) (resp response.IDResp, err error)
}

type NotificationHistoryRepository interface {
	Upsert(ctx context.Context, notificationHistory *entity.NotificationHistory) (err error)
	BatchUpsertTranslate(ctx context.Context, translateNotificationHistories []entity.TranslateNotificationHistory) (
		err error)
	GetPagination(ctx context.Context, req request.GetNotificationPaginationReq) (
		result []entity.GetNotificationHistory, count int64, err error)
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.NotificationHistory, error)
}
