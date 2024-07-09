package contract

import (
	"context"

	"github.com/Halalins/backend/internal/model/response"
	"github.com/google/uuid"
	"mime/multipart"
)

type CloudStorage interface {
	Upload(ctx context.Context, file *multipart.FileHeader, allowedFileTypes []string,
		folder string) (fileName string, err error)
	GenerateSignedURL(ctx context.Context, key string) (url *string, err error)
}

type CloudStorageService interface {
	GenerateSignedURL(ctx context.Context, id uuid.UUID, mediaType string) (response *response.Media, err error)
}
