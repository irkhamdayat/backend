package request

import (
	"bytes"

	"github.com/Halalins/backend/internal/common/constant"
)

type MailerSendEmailReq struct {
	To          string
	Subject     constant.MailerSubject
	BodyContent bytes.Buffer
}

type SendEmailReq struct {
	Template       constant.MailerTemplate
	Subject        constant.MailerSubject
	To             string
	EmailBody      any
	AdditionalData map[string]any
}
