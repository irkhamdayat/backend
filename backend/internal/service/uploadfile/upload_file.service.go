package uploadfile

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/Halalins/backend/internal/model/contract"
)

type Service struct {
	db                     *gorm.DB
	rdb                    *redis.Client
	cloudStorageThirdParty contract.CloudStorage
	uploadFileRepository   contract.UploadFileRepository
	mailerThirdParty       contract.MailerThirdParty
	asynqClient            *asynq.Client
}

func New() *Service {
	return new(Service)
}

func (s *Service) WithPostgresDB(db *gorm.DB) *Service {
	s.db = db
	return s
}

func (s *Service) WithRedisClient(rdb *redis.Client) *Service {
	s.rdb = rdb
	return s
}

func (s *Service) WithCloudStorageThirdParty(thirdParty contract.CloudStorage) *Service {
	s.cloudStorageThirdParty = thirdParty
	return s
}

func (s *Service) WithUploadFileRepository(repo contract.UploadFileRepository) *Service {
	s.uploadFileRepository = repo
	return s
}

func (s *Service) WithMailerThirdParty(thirdParty contract.MailerThirdParty) *Service {
	s.mailerThirdParty = thirdParty
	return s
}

func (s *Service) WithAsynqClient(client *asynq.Client) *Service {
	s.asynqClient = client
	return s
}
