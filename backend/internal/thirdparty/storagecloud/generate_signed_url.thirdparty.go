package storagecloud

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (tp *ThirdParty) GenerateSignedURL(ctx context.Context, key string) (urlPtr *string, err error) {
	var (
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"key": key,
		})
	)

	url, err := tp.storageClient.Bucket(config.Env.GCP.Bucket).SignedURL(key,
		&storage.SignedURLOptions{
			Scheme:  storage.SigningSchemeV4,
			Method:  "GET",
			Expires: time.Now().Add(config.Env.GCP.SignedExpiration),
		},
	)
	if err != nil {
		logger.Errorf("failed generate signed url: %v", err)
		return
	}

	urlPtr = &url

	return
}
