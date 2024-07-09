package notification

import (
	"github.com/Halalins/backend/internal/model/contract"
)

type Handler struct {
	notificationService contract.NotificationService
}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) WithNotificationService(svc contract.NotificationService) *Handler {
	h.notificationService = svc
	return h
}
