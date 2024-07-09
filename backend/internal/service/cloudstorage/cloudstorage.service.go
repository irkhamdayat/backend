package cloudstorage

import (
	"github.com/Halalins/backend/internal/model/contract"
)

type Service struct {
	cloudStorageThirdParty contract.CloudStorage
	uploadFileRepository   contract.UploadFileRepository
}

func New() *Service {
	return new(Service)
}

func (s *Service) WithCloudStorageThirdParty(thirdParty contract.CloudStorage) *Service {
	s.cloudStorageThirdParty = thirdParty
	return s
}

func (s *Service) WithUploadFileRepository(repository contract.UploadFileRepository) *Service {
	s.uploadFileRepository = repository
	return s
}
