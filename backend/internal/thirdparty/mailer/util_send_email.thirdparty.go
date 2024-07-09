package mailer

import (
	"context"
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (rq *ThirdParty) sendEmail(ctx context.Context, req request.MailerSendEmailReq) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
		"req": util.Dump(req),
	})

	m := gomail.NewMessage()
	m.SetHeader("From", config.Env.Mailer.Username)
	m.SetHeader("To", req.To)
	m.SetHeader("Subject", string(req.Subject))
	m.SetBody("text/html", req.BodyContent.String())

	err := rq.gomailDialer.DialAndSend(m)
	if err != nil {
		logger.Errorf("failed dial and send gomaildialer: %v", err)
		return err
	}

	logrus.Info(fmt.Sprintf("Email sent successfully to: %s", req.To))

	return nil
}
