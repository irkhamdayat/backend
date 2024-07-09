package asynq

import (
	"github.com/hibiken/asynq"

	"github.com/Halalins/backend/internal/handler/asynq/mailer"
	"github.com/Halalins/backend/internal/handler/asynq/notification"
	"github.com/Halalins/backend/internal/model/task"
)

type ServeMuxBuilder struct {
	mux *asynq.ServeMux
}

func NewServeMuxBuilder() *ServeMuxBuilder {
	return &ServeMuxBuilder{
		mux: asynq.NewServeMux(),
	}
}

func (b *ServeMuxBuilder) WithMailerHandler(mailerHandler *mailer.Handler) *ServeMuxBuilder {
	b.mux.HandleFunc(task.AsynqSendEmailBoilerplateTask, mailerHandler.SendEmail)
	return b
}

func (b *ServeMuxBuilder) WithNotificationService(notification *notification.Handler) *ServeMuxBuilder {
	b.mux.HandleFunc(task.AsynqCreateNotification, notification.CreateNotification)
	return b
}

func (b *ServeMuxBuilder) WithMiddleware(mid asynq.MiddlewareFunc) *ServeMuxBuilder {
	b.mux.Use(mid)
	return b
}

func (b *ServeMuxBuilder) Build() *asynq.ServeMux {
	return b.mux
}
