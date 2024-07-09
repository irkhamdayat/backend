package mailer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (rq *ThirdParty) Send(ctx context.Context, req request.SendEmailReq) error {
	var (
		body   = map[string]any{}
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"req": util.Dump(req),
		})
	)

	tmp, _ := json.Marshal(req.EmailBody)
	_ = json.Unmarshal(tmp, &body)
	body["Year"] = time.Now().Year()

	bodyContent, err := rq.processMailTemplate(req.Template, body)
	if err != nil {
		logger.Errorf("failed process mail template: %v", err)
		return err
	}

	err = rq.sendEmail(ctx, request.MailerSendEmailReq{
		To:          req.To,
		Subject:     req.Subject,
		BodyContent: *bodyContent,
	})
	if err != nil {
		logger.Errorf("failed send email: %v", err)
		return err
	}

	return nil
}
