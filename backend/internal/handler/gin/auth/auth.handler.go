package auth

//
import "github.com/Halalins/backend/internal/model/contract"

type Handler struct {
	adminService contract.AdminService
	agentService contract.AgentService
}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) WithAdminService(svc contract.AdminService) *Handler {
	h.adminService = svc
	return h
}

func (h *Handler) WithAgentService(svc contract.AgentService) *Handler {
	h.agentService = svc
	return h
}
