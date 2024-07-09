package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Halalins/backend/internal/model/response"
)

func (h *Handler) LoginOrRefreshResponse(c *gin.Context, _ int, token string, expiredAt time.Time) {
	c.JSON(http.StatusOK, response.LoginResponse{
		ExpiredAt:   expiredAt.Format(time.RFC3339),
		AccessToken: token,
	})
}
