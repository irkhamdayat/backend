package agent

import (
	"context"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Service) GetAgentInfo(ctx context.Context, id uuid.UUID) (*response.GetAgentInfoResp, error) {
	var (
		agent  *entity.Agent
		logger = logrus.WithFields(logrus.Fields{
			"ctx": util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"id":  id,
		})
	)

	agent, err := s.agentRepository.FindByID(ctx, id)
	if err != nil {
		logger.Errorf("failed get agent: %v", err)
		return nil, err
	}

	resp := &response.GetAgentInfoResp{
		ID:                agent.ID,
		Code:              agent.Code,
		PhoneNumber:       agent.PhoneNumber,
		BirthDate:         agent.BirthDate.Format("02-01-2006"),
		BirthPlace:        agent.BirthPlace,
		Address:           agent.Address,
		Location:          agent.Location,
		FirstName:         agent.FirstName,
		LastName:          agent.LastName,
		Email:             agent.Email,
		Status:            agent.Status,
		CodeReferral:      agent.CodeReferral,
		KtpNumber:         agent.KtpNumber,
		NpwpNumber:        agent.NpwpNumber,
		BankAccountNumber: agent.BankAccountNumber,
		Bank: response.Bank{
			ID:   agent.Bank.ID,
			Code: agent.Bank.Code,
		},
		IsSubscribeNews: agent.IsSubscribeNews,
	}

	return resp, nil
}
