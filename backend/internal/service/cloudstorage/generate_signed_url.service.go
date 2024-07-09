package cloudstorage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/response"
)

func (s *Service) GenerateSignedURL(ctx context.Context, id uuid.UUID, mediaType string) (resp *response.Media, err error) {
	var (
		url          *string
		uploadedFile *entity.UploadFile
		logger       = logrus.WithFields(logrus.Fields{
			"id":        id,
			"mediaType": mediaType,
		})
	)

	uploadedFile, err = s.uploadFileRepository.FindByIDAndUploadType(ctx, id,
		mediaType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = constant.ErrDataNotFound
		return nil, err
	}

	if err != nil {
		logger.Errorf("failed get upload file: %v", err)
		return nil, err
	}

	url, err = s.cloudStorageThirdParty.GenerateSignedURL(ctx, uploadedFile.FileName)
	if err != nil {
		logger.Errorf("failed generate signed url: %v", err)
		return nil, err
	}

	resp = &response.Media{
		ID:  uploadedFile.ID,
		Url: *url,
	}

	return
}
