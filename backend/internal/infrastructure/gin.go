package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
)

func InitializeGinServer() *http.Server {
	switch config.Env.Env {
	case constant.EnvDevelopment, constant.EnvStaging:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	util.AddValidation()

	return &http.Server{
		Addr:           ":" + config.Env.App.Port,
		ReadTimeout:    time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
}
