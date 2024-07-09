package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (h *Handler) Authenticator(userType string) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var (
			ctx    = c.Request.Context()
			logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
		)

		switch userType {
		case constant.AdminUserType:
			var req request.AdminLoginReq
			if err := c.ShouldBindJSON(&req); err != nil {
				c.Set("should-use-err", err)
				errmapper.HandleError(c, err)
				return nil, err
			}
			userClaim, err := h.adminService.PostLogin(ctx, req)
			if err != nil {
				c.Set("should-use-err", err)
				logger.Errorf("failed login: %v", err)
				errmapper.HandleError(c, err)
				return nil, err
			}
			return userClaim, nil

		case constant.AgentUserType:
			var req request.AgentLoginReq
			if err := c.ShouldBindJSON(&req); err != nil {
				c.Set("should-use-err", err)
				errmapper.HandleError(c, err)
				return nil, err
			}
			userClaim, err := h.agentService.PostLogin(ctx, req)
			if err != nil {
				c.Set("should-use-err", err)
				logger.Errorf("failed login: %v", err)
				errmapper.HandleError(c, err)
				return nil, err
			}
			return userClaim, nil

		default:
			err := fmt.Errorf("unknown user type")
			c.Set("should-use-err", err)
			errmapper.HandleError(c, err)
			return nil, err
		}
	}
}
