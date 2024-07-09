package admin

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByUsernameAndEncryptedPassword(ctx context.Context, username, encryptedPassword string) (
	admin *entity.Admin, err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":      util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"username": username,
		})
		db = util.GetTxFromContext(ctx, r.db)
	)

	err = db.WithContext(ctx).
		Where("username = ?", username).
		Where("password = ?", encryptedPassword).
		First(&admin).
		Error

	if err != nil {
		logger.Errorf("failed to get admin: %v", err)
		return nil, err
	}

	return admin, nil
}
