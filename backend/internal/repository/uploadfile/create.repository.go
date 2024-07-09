package uploadfile

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) Create(ctx context.Context, uploadFile entity.UploadFile) (*uuid.UUID, error) {
	var (
		tx     = util.GetTxFromContext(ctx, r.db)
		logger = logrus.WithFields(logrus.Fields{
			"ctx":    util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"entity": util.Dump(uploadFile),
		})
	)

	err := tx.WithContext(ctx).
		Create(&uploadFile).
		Error

	if err != nil {
		logger.Errorf("failed create upload file: %v", err)
		return nil, err
	}

	return &uploadFile.ID, nil
}
