package admin

import (
	"github.com/Halalins/backend/internal/model/contract"
)

type Handler struct {
	adminService contract.AdminService
}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) WithAdminService(svc contract.AdminService) *Handler {
	h.adminService = svc
	return h
}
