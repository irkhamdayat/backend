package admin

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/Halalins/backend/internal/model/contract"
)

type Service struct {
	db                   *gorm.DB
	rdb                  *redis.Client
	adminRepository      contract.AdminRepository
	asynqClient          *asynq.Client
	storageService       contract.CloudStorageService
	uploadFileRepository contract.UploadFileRepository
	//roleRepository         contract.RoleRepository
	//roleToAccessRepository contract.RoleToAccessRepository
	//masterDataRepository   contract.MasterDataRepository
}

func New() *Service {
	return new(Service)
}

func (s *Service) WithRedisClient(rdb *redis.Client) *Service {
	s.rdb = rdb
	return s
}

func (s *Service) WithPostgresDB(db *gorm.DB) *Service {
	s.db = db
	return s
}

func (s *Service) WithAdminRepository(repository contract.AdminRepository) *Service {
	s.adminRepository = repository
	return s
}

func (s *Service) WithAsynqClient(client *asynq.Client) *Service {
	s.asynqClient = client
	return s
}

func (s *Service) WithCloudStorageService(svc contract.CloudStorageService) *Service {
	s.storageService = svc
	return s
}

func (s *Service) WithUploadFileRepository(repository contract.UploadFileRepository) *Service {
	s.uploadFileRepository = repository
	return s
}

//func (s *Service) WithRoleRepository(repo contract.RoleRepository) *Service {
//	s.roleRepository = repo
//	return s
//}
//
//func (s *Service) WithRoleToAccessRepository(repo contract.RoleToAccessRepository) *Service {
//	s.roleToAccessRepository = repo
//	return s
//}
//
//func (s *Service) WithMasterDataRepository(repo contract.MasterDataRepository) *Service {
//	s.masterDataRepository = repo
//	return s
//}
