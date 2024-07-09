package uploadfile

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (s *Service) GetSignedURLFile(ctx context.Context, req request.GetSignedURLFileReq) (url *string, err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx":        util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"ID":         req.ID.String(),
			"uploadType": req.UploadType,
		})
	)

	uploadFile, err := s.uploadFileRepository.FindByIDAndUploadType(ctx, req.ID, req.UploadType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = constant.ErrDataNotFound
		return
	}

	if err != nil {
		logger.Errorf("failed get upload file: %v", err)
		return
	}

	cacheKey := cachekey.NewGetSignedURLCacheKey(uploadFile.FileName)
	url, err = util.GetOrSetCache[string](ctx, s.rdb, cacheKey, config.Env.GCP.SignedExpiration, func() (*string,
		error) {
		url, err = s.cloudStorageThirdParty.GenerateSignedURL(ctx, uploadFile.FileName)
		if err != nil {
			logger.Errorf("failed get generate signed url file: %v", err)
			return nil, err
		}

		return url, nil
	})
	if err != nil {
		logger.Errorf("failed get get or set cache: %v", err)
		return
	}

	return
}
