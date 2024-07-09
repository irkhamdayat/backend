package agent

import (
	"github.com/Halalins/backend/internal/model/contract"
)

type Handler struct {
	agentService contract.AgentService
}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) WithAgentService(svc contract.AgentService) *Handler {
	h.agentService = svc
	return h
}
