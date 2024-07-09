package infrastructure

import (
	"github.com/go-gomail/gomail"

	"github.com/Halalins/backend/config"
)

func InitializeGomailDialer() *gomail.Dialer {
	return gomail.NewDialer(
		config.Env.Mailer.Host,
		config.Env.Mailer.Port,
		config.Env.Mailer.Username,
		config.Env.Mailer.Password,
	)
}
