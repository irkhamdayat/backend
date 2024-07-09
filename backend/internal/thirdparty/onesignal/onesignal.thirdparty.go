package onesignal

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/OneSignal/onesignal-go-api"
)

type ThirdParty struct {
	onesignalClient *onesignal.APIClient
	i18nBundle      *i18n.Bundle
}

func New() *ThirdParty {
	configuration := onesignal.NewConfiguration()

	return &ThirdParty{
		onesignalClient: onesignal.NewAPIClient(configuration),
	}
}

func (s *ThirdParty) WithI18nBundle(bundle *i18n.Bundle) *ThirdParty {
	s.i18nBundle = bundle
	return s
}
