package cachekey

import (
	"fmt"

	"github.com/google/uuid"
)

func NewGetUploadFileCacheKey(ID uuid.UUID, uploadType string) string {
	return fmt.Sprintf("upload-file:id:%s:upload-type:%s", ID.String(), uploadType)
}

func NewGetSignedURLCacheKey(fileName string) string {
	return fmt.Sprintf("signed-url:filename:%s", fileName)
}
