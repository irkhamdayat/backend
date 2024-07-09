package mailer

import "github.com/Halalins/backend/internal/model/contract"

type Handler struct {
	mailerService contract.MailerService
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) WithMailerService(service contract.MailerService) *Handler {
	s.mailerService = service
	return s
}
