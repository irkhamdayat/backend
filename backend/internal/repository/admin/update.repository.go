package admin

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) Update(ctx context.Context, admin *entity.Admin) (err error) {
	var (
		tx     = util.GetTxFromContext(ctx, r.db)
		logger = logrus.WithFields(logrus.Fields{
			"ctx":   util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"admin": util.Dump(admin),
		})
	)

	err = tx.WithContext(ctx).
		Updates(&admin).
		Error

	if err != nil {
		logger.Errorf("failed to update admin: %v", err)
		return
	}

	return
}
