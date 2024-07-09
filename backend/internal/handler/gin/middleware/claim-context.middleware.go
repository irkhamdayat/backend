package middleware

import (
	"context"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/contract"

	"github.com/gin-gonic/gin"
)

type ClaimContext struct {
	adminRepository contract.AdminRepository
	agentRepository contract.AgentRepository
}

func NewAccountSetupMiddleware() *ClaimContext {
	return new(ClaimContext)
}

func (m *ClaimContext) WithAdminRepository(repository contract.AdminRepository) *ClaimContext {
	m.adminRepository = repository
	return m
}

func (m *ClaimContext) WithAgentRepository(repository contract.AgentRepository) *ClaimContext {
	m.agentRepository = repository
	return m
}

func (m *ClaimContext) SetupAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id, err := util.GetUserIDFromContext(ctx, constant.AdminUserType)
		if err != nil {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		admin, err := m.adminRepository.FindByID(ctx, *id)
		if err != nil {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		if admin != nil {
			reqCtx := context.WithValue(ctx, constant.RoleIDKey, admin.Role.ID)                 //nolint
			reqCtx = context.WithValue(ctx, constant.BrandInsuranceKey, admin.InsuranceBrandID) //nolint
			c.Request = c.Request.WithContext(reqCtx)
		}

		c.Next()
	}
}

func (m *ClaimContext) SetupAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id, err := util.GetUserIDFromContext(ctx, constant.AgentUserType)
		if err != nil {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		_, err = m.agentRepository.FindByID(ctx, *id)
		if err != nil {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
