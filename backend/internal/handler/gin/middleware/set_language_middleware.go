package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Halalins/backend/internal/common/constant"
)

func SetLanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		acceptLang = strings.ToUpper(acceptLang)

		if !constant.AcceptLanguage[acceptLang] {
			acceptLang = constant.LangDefault
		}

		c.Set(constant.Lang, acceptLang)

		//nolint
		reqCtx := context.WithValue(c.Request.Context(), constant.Lang, acceptLang)
		c.Request = c.Request.WithContext(reqCtx)

		c.Next()
	}
}
