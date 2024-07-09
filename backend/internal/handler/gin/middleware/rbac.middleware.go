package middleware

import (
	"errors"
	"github.com/Halalins/backend/internal/common/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/model/contract"
	"github.com/Halalins/backend/internal/model/entity"
)

type RBACMiddleware struct {
	roleToAccessRepository contract.RoleToAccessRepository
	adminRepository        contract.AdminRepository
}

func NewRBACMiddleware() *RBACMiddleware {
	return new(RBACMiddleware)
}

func (m *RBACMiddleware) WithRoleToAccessRepository(repo contract.RoleToAccessRepository) *RBACMiddleware {
	m.roleToAccessRepository = repo
	return m
}

func (m *RBACMiddleware) WithAdminRepository(repo contract.AdminRepository) *RBACMiddleware {
	m.adminRepository = repo
	return m
}

func (m *RBACMiddleware) ValidateAccess(accessID uuid.UUID) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx = c.Request.Context()
		)
		userClaim, err := util.BindingFromContext[entity.UserClaim](ctx, constant.ListUserClaimKeys, entity.ValidatorEntityAdmin)
		if err != nil {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		admin, err := m.adminRepository.FindByID(ctx, userClaim.ID)
		if err != nil {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		roleToAccess, err := m.roleToAccessRepository.FindByRoleIDAndAccessID(ctx, admin.RoleID, accessID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		if roleToAccess == nil {
			errmapper.HandleError(c, constant.ErrAccessPermissionDenied)
			c.Abort()
			return
		}

		c.Next()
	}
}
