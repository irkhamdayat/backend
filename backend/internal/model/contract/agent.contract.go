package contract

import (
	"context"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/google/uuid"

	"github.com/Halalins/backend/internal/model/entity"
)

type AgentRepository interface {
	FindByUsernameOrEmailAndPassword(ctx context.Context, username, email, encryptedPassword string) (*entity.Agent, error)
	FindByEmailOrUsername(ctx context.Context, email, username string) (*entity.Agent, error)
	Create(ctx context.Context, admin *entity.Agent) error
	Update(ctx context.Context, companyUser *entity.Agent) error
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.Agent, error)
	FindByIDAndPin(ctx context.Context, id uuid.UUID, encryptedPin string) (*entity.Agent, error)
}

type AgentService interface {
	PostLogin(ctx context.Context, req request.AgentLoginReq) (*entity.UserClaim, error)
	GetAgentInfo(ctx context.Context, id uuid.UUID) (*response.GetAgentInfoResp, error)
	PostRegister(ctx context.Context, req request.AgentRegisterReq) (resp *response.IDResp, err error)
	PostVerifyEmail(ctx context.Context, req request.PostVerifyEmailReq) (*response.PostVerifyEmailResp, error)
	PostVerifyPin(ctx context.Context, req request.PostVerifyPinReq) (*response.PostVerifyPinResp, error)
}
