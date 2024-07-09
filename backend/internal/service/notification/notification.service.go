package notification

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"

	"github.com/Halalins/backend/internal/model/contract"
	"github.com/Halalins/backend/internal/thirdparty/onesignal"
)

type Service struct {
	onesignalThirdParty           *onesignal.ThirdParty
	notificationHistoryRepository contract.NotificationHistoryRepository
	db                            *gorm.DB
	i18nBundle                    *i18n.Bundle
}

func New() *Service {
	return new(Service)
}

func (s *Service) WithPostgresDB(db *gorm.DB) *Service {
	s.db = db
	return s
}

func (s *Service) WithOnesignalRequester(thirdParty *onesignal.ThirdParty) *Service {
	s.onesignalThirdParty = thirdParty
	return s
}

func (s *Service) WithNotificationHistoryRepository(repository contract.NotificationHistoryRepository) *Service {
	s.notificationHistoryRepository = repository
	return s
}

func (s *Service) WithI18nBundle(bundle *i18n.Bundle) *Service {
	s.i18nBundle = bundle
	return s
}
