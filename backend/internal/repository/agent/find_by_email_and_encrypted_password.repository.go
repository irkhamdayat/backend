package agent

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByUsernameOrEmailAndPassword(ctx context.Context, username, email, encryptedPassword string) (
	agent *entity.Agent, err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"username": username,
		})
		db = util.GetTxFromContext(ctx, r.db)
	)

	err = db.WithContext(ctx).
		Where("username = ? OR email = ?", username, email).
		Where("password = ?", encryptedPassword).
		First(&agent).
		Error

	if err != nil {
		logger.Errorf("failed to get agent: %v", err)
		return nil, err
	}

	return agent, nil
}
