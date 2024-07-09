package mailer

import (
	"github.com/go-gomail/gomail"
)

type ThirdParty struct {
	gomailDialer *gomail.Dialer
}

func New() *ThirdParty {
	return new(ThirdParty)
}

func (rq *ThirdParty) WithGomailDialer(dialer *gomail.Dialer) *ThirdParty {
	rq.gomailDialer = dialer
	return rq
}
