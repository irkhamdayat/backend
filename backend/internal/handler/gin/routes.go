package gin

import (
	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/handler/gin/admin"
	"github.com/Halalins/backend/internal/handler/gin/agent"
	"github.com/Halalins/backend/internal/handler/gin/healthcheck"
	"github.com/Halalins/backend/internal/handler/gin/middleware"
	"github.com/Halalins/backend/internal/handler/gin/uploadfile"
	jwt "github.com/appleboy/gin-jwt/v2"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	healthCheckHandler *healthcheck.Handler,
	uploadFileHandler *uploadfile.Handler,
	adminHandler *admin.Handler,
	adminAuthMiddleware *jwt.GinJWTMiddleware,
	claimContextMiddleware *middleware.ClaimContext,
	rbacMiddleware *middleware.RBACMiddleware,
	agentHandler *agent.Handler,
	agentAuthMiddleware *jwt.GinJWTMiddleware,
	rateLimitMiddleware *middleware.RateLimitMiddleware,
) *gin.Engine {
	r := gin.New()

	// sentry apm middleware
	if config.Env.Sentry.EnableAPM {
		r.Use(sentrygin.New(sentrygin.Options{}))
	}

	// using cors
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           nil,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.SetLanguageMiddleware())

	//SCOPE: Public URL
	v1Group := r.Group("/v1")
	v1Group.Any("/ping", healthCheckHandler.Ping)

	mediaGroup := v1Group.Group("/medias")
	mediaGroup.GET(":id", uploadFileHandler.RedirectToSignedURL)
	mediaGroup.POST("", uploadFileHandler.UploadFile)

	//SCOPE: Agent Region
	agentGroup := v1Group.Group("/agents")

	agentAuthGroup := agentGroup.Group("/auth")
	agentAuthGroup.POST("/register", agentHandler.PostRegister)
	agentAuthGroup.POST("/verify-email", agentHandler.PostVerifyEmail)
	agentAuthGroup.POST("/login", agentAuthMiddleware.LoginHandler)
	agentAuthGroup.POST("/refresh", agentAuthMiddleware.RefreshHandler)

	agentGroup.Use(adminAuthMiddleware.MiddlewareFunc())
	agentGroup.Use(claimContextMiddleware.SetupAgentMiddleware())

	agentGroup.POST("verify-pin", rateLimitMiddleware.Rate("verify-pin"), agentHandler.PostVerifyPin)

	//SCOPE: Admin Region
	adminGroup := v1Group.Group("/admins")

	adminAuthGroup := adminGroup.Group("/auth")
	adminAuthGroup.POST("/register", adminHandler.PostRegister)
	adminAuthGroup.POST("/forgot", adminHandler.PostForgotPassword)
	adminAuthGroup.POST("/change-password", adminHandler.PostChangePassword)
	adminAuthGroup.POST("/login", adminAuthMiddleware.LoginHandler)
	adminAuthGroup.POST("/refresh", adminAuthMiddleware.RefreshHandler)

	adminGroup.Use(adminAuthMiddleware.MiddlewareFunc())
	adminGroup.Use(claimContextMiddleware.SetupAdminMiddleware())

	adminGroup.GET("", adminHandler.GetAdminInfo)

	return r
}
