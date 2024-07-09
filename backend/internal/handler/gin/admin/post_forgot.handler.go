package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (h *Handler) PostForgotPassword(c *gin.Context) {
	var (
		req    request.AdminForgotPasswordReq
		ctx    = c.Request.Context()
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	if err := c.ShouldBind(&req); err != nil {
		logger.Errorf("failed binding request: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	resp, err := h.adminService.PostForgotPassword(ctx, req)
	if err != nil {
		logger.Errorf("failed call forgot password service: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
