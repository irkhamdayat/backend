package uploadfile

import "github.com/Halalins/backend/internal/model/contract"

type Handler struct {
	uploadFileService contract.UploadFileService
}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) WithUploadFileService(svc contract.UploadFileService) *Handler {
	h.uploadFileService = svc
	return h
}
