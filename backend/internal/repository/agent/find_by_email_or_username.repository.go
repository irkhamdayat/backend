package agent

import (
	"context"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"

	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByEmailOrUsername(ctx context.Context, email, username string) (agent *entity.Agent,
	err error) {
	var (
		db     = util.GetTxFromContext(ctx, r.db)
		logger = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"email":    email,
			"username": username,
		})
	)

	err = db.WithContext(ctx).
		Where("email = ? OR username = ?", email, username).
		First(&agent).
		Error

	if err != nil {
		logger.Errorf("failed to get agent: %v", err)
		return
	}

	return
}
