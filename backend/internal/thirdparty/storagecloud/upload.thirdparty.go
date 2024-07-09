package storagecloud

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"slices"

	mimetype "github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
)

func (tp *ThirdParty) Upload(ctx context.Context, file *multipart.FileHeader, allowedFileTypes []string,
	folder string) (fileName string, err error) {
	var (
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	fileOpened, err := file.Open()
	if err != nil {
		logger.Errorf("failed open file: %v", err)
		return
	}

	defer func() {
		_ = fileOpened.Close()
	}()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, fileOpened)
	if err != nil {
		logger.Errorf("failed copy file: %v", err)
		return
	}

	if buf.Len() == 0 {
		logger.Error("Buffer is empty after copying data")
		return
	}

	mimeType := mimetype.Detect(buf.Bytes())
	fileExtension := mimeType.Extension()
	if !slices.Contains(allowedFileTypes, fileExtension) {
		logger.Error(err)
		err = constant.ErrFileTypeIsNotSupported
		return
	}

	fileName = fmt.Sprintf("%s/%s-%s%s",
		folder,
		util.GenerateRandomString(6, constant.AlphaNumeric),
		uuid.New().String(),
		fileExtension,
	)

	wc := tp.storageClient.Bucket(config.Env.GCP.Bucket).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, buf); err != nil {
		logger.Errorf("failed io.Copy: %v", err)
		return
	}

	if err = wc.Close(); err != nil {
		logger.Errorf("failed Writer.Close: %v", err)
		return
	}

	return
}
