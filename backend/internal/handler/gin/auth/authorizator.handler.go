package auth

import "github.com/gin-gonic/gin"

func (h *Handler) Authorizator(_ interface{}, _ *gin.Context) bool {
	return true
}
