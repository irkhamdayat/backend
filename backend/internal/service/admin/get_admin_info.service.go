package admin

import (
	"context"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

func (s *Service) GetAdminInfo(ctx context.Context, ID uuid.UUID) (*response.GetAdminInfoResp, error) {
	var (
		admin  *entity.Admin
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"id":  ID,
		})
	)

	admin, err := s.adminRepository.FindByID(ctx, ID)
	if err != nil {
		logger.Errorf("failed get admin: %v", err)
		return nil, err
	}

	resp := &response.GetAdminInfoResp{
		ID:        admin.ID,
		FirstName: admin.FirstName,
		LastName:  null.NewString(admin.LastName.String, admin.LastName.Valid),
		Username:  admin.Username,
		Email:     admin.Email,
		Role: response.Role{
			ID:   admin.Role.ID,
			Name: admin.Role.Name,
		},
		//TODO: Fix this response, should be insurance with insurance response
		InsuranceBrandID: null.NewString(admin.InsuranceBrandID.String, admin.InsuranceBrandID.Valid),
		Status:           admin.Status,
	}

	if admin.Photo.Valid {
		var media *response.Media
		media, err = s.storageService.GenerateSignedURL(ctx, admin.Photo.UUID, constant.UploadTypeProfilePicture)
		if err != nil {
			logger.Errorf("failed get signed url: %v", err)
			return nil, err
		}

		resp.Photo = media
	}

	return resp, nil
}
