package contract

import (
	"context"

	"github.com/google/uuid"

	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
)

type UploadFileRepository interface {
	Create(ctx context.Context, entity entity.UploadFile) (*uuid.UUID, error)
	FindByIDAndUploadType(ctx context.Context, ID uuid.UUID, uploadType string) (*entity.UploadFile, error)
}

type UploadFileService interface {
	Upload(ctx context.Context, req request.UploadFileReq) (*response.IDResp, error)
	GetSignedURLFile(ctx context.Context, req request.GetSignedURLFileReq) (*string, error)
}
