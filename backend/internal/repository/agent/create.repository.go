package agent

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) Create(ctx context.Context, agent *entity.Agent) (err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":   util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"admin": agent,
		})
		db = util.GetTxFromContext(ctx, r.db)
	)

	err = db.WithContext(ctx).
		Create(&agent).
		Error

	if err != nil {
		logger.Errorf("failed create agent: %v", err)
		return
	}

	return
}
