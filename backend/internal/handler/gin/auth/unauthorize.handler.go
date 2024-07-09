package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
)

func (h *Handler) Unauthorized(c *gin.Context, _ int, _ string) {
	if _, ok := c.Get("should-use-err"); ok {
		return
	}
	errmapper.HandleError(c, constant.ErrUnauthorized)
}
