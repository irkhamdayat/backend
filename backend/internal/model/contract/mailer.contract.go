package contract

import (
	"context"

	"github.com/Halalins/backend/internal/model/request"
)

type MailerThirdParty interface {
	Send(ctx context.Context, req request.SendEmailReq) error
}

type MailerService interface {
	Send(ctx context.Context, req request.SendEmailReq) error
}
