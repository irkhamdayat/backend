package admin

import (
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetAdminInfo(c *gin.Context) {
	var (
		req    request.GetAdminInfoReq
		ctx    = c.Request.Context()
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	if err := c.ShouldBind(&req); err != nil {
		logger.Errorf("failed binding request: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	id, err := util.GetUserIDFromContext(ctx, constant.EntityTypeAdmin)
	if err != nil {
		return
	}

	req.ID = *id

	resp, err := h.adminService.GetAdminInfo(ctx, req.ID)
	if err != nil {
		logger.Errorf("failed call get admin info service: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
