package contract

import (
	"context"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/google/uuid"

	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

type AdminRepository interface {
	FindByUsernameAndEncryptedPassword(ctx context.Context, username, encryptedPassword string) (*entity.Admin, error)
	FindByEmailOrUsername(ctx context.Context, email, username string) (*entity.Admin, error)
	Create(ctx context.Context, admin *entity.Admin) error
	Update(ctx context.Context, companyUser *entity.Admin) error
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.Admin, error)
}

type AdminService interface {
	PostLogin(ctx context.Context, req request.AdminLoginReq) (*entity.UserClaim, error)
	PostChangePassword(ctx context.Context, req request.AdminChangePasswordReq) (*response.IDResp, error)
	PostRegister(ctx context.Context, req request.AdminRegisterReq) (resp *response.IDResp, err error)
	PostForgotPassword(ctx context.Context, req request.AdminForgotPasswordReq) (*response.EmailResp, error)
	PostAddAdmin(ctx context.Context, req request.AdminRegisterReq) (resp *response.IDResp, err error)
	GetAdminInfo(ctx context.Context, ID uuid.UUID) (*response.GetAdminInfoResp, error)
}
