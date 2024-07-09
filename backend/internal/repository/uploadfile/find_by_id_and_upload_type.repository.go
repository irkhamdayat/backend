package uploadfile

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/Halalins/backend/internal/model/entity"
)

func (r *Repository) FindByIDAndUploadType(ctx context.Context, ID uuid.UUID, uploadType string) (*entity.UploadFile, error) {
	var (
		cacheKey   = cachekey.NewGetUploadFileCacheKey(ID, uploadType)
		uploadFile *entity.UploadFile
		logger     = logrus.WithFields(logrus.Fields{
			"ctx":        util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"uploadType": uploadType,
		})
	)

	uploadFile, err := util.GetOrSetCache[entity.UploadFile](ctx, r.rdb, cacheKey, config.Env.GCP.SignedExpiration,
		func() (*entity.UploadFile, error) {
			err := r.db.WithContext(ctx).
				Where("id = ? AND upload_type = ?", ID, uploadType).
				First(&uploadFile).
				Error

			if err != nil {
				logger.Errorf("failed get upload file: %v", err)
				return nil, err
			}

			return uploadFile, nil
		})
	if err != nil {
		logger.Errorf("failed get or set cache: %v", err)
		return nil, err
	}

	return uploadFile, nil
}
