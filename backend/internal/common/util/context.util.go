package util

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Halalins/backend/internal/common/constant"
)

func GetAcceptLangFromGinContext(ctx *gin.Context) string {
	lang := ctx.Keys[constant.Lang]
	if lang == nil {
		return constant.LangDefault
	}
	return fmt.Sprintf("%s", lang)
}

func GetAcceptLanguageFromContext(ctx context.Context) string {
	lang := ctx.Value(constant.Lang)
	if lang == nil {
		lang = constant.LangDefault
	}
	return fmt.Sprintf("%s", lang)
}
