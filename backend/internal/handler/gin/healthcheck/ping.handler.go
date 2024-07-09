package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Halalins/backend/internal/model/response"
)

func (c *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.MessageResponse{Message: "Pong!!!"})
}
