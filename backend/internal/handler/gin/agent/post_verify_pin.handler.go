package agent

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

func (h *Handler) PostVerifyPin(c *gin.Context) {
	var (
		req    request.PostVerifyPinReq
		ctx    = c.Request.Context()
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	if err := c.ShouldBind(&req); err != nil {
		logger.Errorf("failed binding request: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	id, err := util.GetUserIDFromContext(ctx, constant.EntityTypeAgent)
	if err != nil {
		return
	}

	req.ID = *id

	resp, err := h.agentService.PostVerifyPin(ctx, req)
	if err != nil {
		logger.Errorf("failed call service: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
