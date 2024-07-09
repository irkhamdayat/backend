package mailer

import (
	"github.com/Halalins/backend/internal/thirdparty/mailer"
)

type Service struct {
	mailerThirdParty *mailer.ThirdParty
}

func New() *Service {
	return new(Service)
}

func (s *Service) WithMailerThirdParty(thirdParty *mailer.ThirdParty) *Service {
	s.mailerThirdParty = thirdParty
	return s
}
