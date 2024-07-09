package uploadfile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (h *Handler) RedirectToSignedURL(c *gin.Context) {
	var (
		req    request.RedirectToS3SignedURLReq
		ctx    = c.Request.Context()
		logger = logrus.WithField("ctx", util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys))
	)

	if err := c.ShouldBindUri(&req); err != nil {
		logger.Errorf("failed binding request: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	resp, err := h.uploadFileService.GetSignedURLFile(ctx, request.GetSignedURLFileReq{
		ID:         uuid.MustParse(req.ID),
		UploadType: constant.UploadTypeRichMedia,
	})
	if err != nil {
		logger.Errorf("failed get signed url file service: %v", err)
		errmapper.HandleError(c, err)
		return
	}

	c.Redirect(http.StatusPermanentRedirect, *resp)
}
