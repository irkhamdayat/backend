package bootstrap

import (
	//Prebuild Library
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"

	//External Library
	"cloud.google.com/go/storage"
	"github.com/common-nighthawk/go-figure"
	"github.com/go-gomail/gomail"
	"github.com/hibiken/asynq"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"

	//Internal Library: Common
	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/infrastructure"

	//Internal Library: Thirdparty
	"github.com/Halalins/backend/internal/thirdparty/mailer"
	"github.com/Halalins/backend/internal/thirdparty/storagecloud"

	//Internal Library: Repository
	adminRepo "github.com/Halalins/backend/internal/repository/admin"
	agentRepo "github.com/Halalins/backend/internal/repository/agent"
	roleToAccessRepo "github.com/Halalins/backend/internal/repository/roletoacess"
	uploadFileRepo "github.com/Halalins/backend/internal/repository/uploadfile"

	//Internal Library: Service
	adminSvc "github.com/Halalins/backend/internal/service/admin"
	agentSvc "github.com/Halalins/backend/internal/service/agent"
	"github.com/Halalins/backend/internal/service/cloudstorage"
	uploadFileSvc "github.com/Halalins/backend/internal/service/uploadfile"

	//Internal Library: Handler
	"github.com/Halalins/backend/internal/handler/gin"
	adminHndlr "github.com/Halalins/backend/internal/handler/gin/admin"
	agentHndlr "github.com/Halalins/backend/internal/handler/gin/agent"
	"github.com/Halalins/backend/internal/handler/gin/auth"
	"github.com/Halalins/backend/internal/handler/gin/healthcheck"
	"github.com/Halalins/backend/internal/handler/gin/middleware"
	uploadFileHndlr "github.com/Halalins/backend/internal/handler/gin/uploadfile"
)

type AppServer struct {
	PostgresDB         *gorm.DB
	RedisClient        *redis.Client
	HTPPGinServer      *http.Server
	AsynqClient        *asynq.Client
	StorageCloudClient *storage.Client
	GomailDialer       *gomail.Dialer
	I18nBundle         *i18n.Bundle
}

func greetingServer() {
	figure.NewColorFigure(config.Env.App.Name, "doom", "red", true).Print()
}

func StartServer() {
	greetingServer()

	infrastructure.InitializeSentry(ServiceName+"-server", ServiceVersion)

	app := AppServer{
		PostgresDB:         infrastructure.InitializePostgresConn(),
		RedisClient:        infrastructure.InitializeRedisConn(),
		HTPPGinServer:      infrastructure.InitializeGinServer(),
		AsynqClient:        infrastructure.InitializeAsynqClient(),
		StorageCloudClient: infrastructure.InitializeStorageCloud(),
		GomailDialer:       infrastructure.InitializeGomailDialer(),
		I18nBundle:         infrastructure.InitializeI18nBundle(),
	}

	_ = errmapper.Initialize().WithI18nBundle(app.I18nBundle, "errors")

	pgDB, err := app.PostgresDB.DB()
	util.ContinueOrFatal(err)

	// init thirdparty
	storageThirdParty := storagecloud.New().
		WithStorageClient(app.StorageCloudClient)
	mailerThirdParty := mailer.New().
		WithGomailDialer(app.GomailDialer)

	// init repositories
	uploadfileRepository := uploadFileRepo.New().
		WithRedisClient(app.RedisClient).
		WithPostgresDB(app.PostgresDB)
	adminRepository := adminRepo.New().
		WithPostgresDB(app.PostgresDB)
	roleToAccessRepository := roleToAccessRepo.New().
		WithPostgresDB(app.PostgresDB)
	agentRepository := agentRepo.New().
		WithPostgresDB(app.PostgresDB)

	// init services
	cloudStorageService := cloudstorage.New().
		WithCloudStorageThirdParty(storageThirdParty).
		WithUploadFileRepository(uploadfileRepository)
	uploadFileService := uploadFileSvc.New().
		WithPostgresDB(app.PostgresDB).
		WithRedisClient(app.RedisClient).
		WithCloudStorageThirdParty(storageThirdParty).
		WithUploadFileRepository(uploadfileRepository).
		WithMailerThirdParty(mailerThirdParty).
		WithAsynqClient(app.AsynqClient)
	adminService := adminSvc.New().
		WithPostgresDB(app.PostgresDB).
		WithRedisClient(app.RedisClient).
		WithAsynqClient(app.AsynqClient).
		WithAdminRepository(adminRepository).
		WithCloudStorageService(cloudStorageService).
		WithUploadFileRepository(uploadfileRepository)
	agentService := agentSvc.New().
		WithPostgresDB(app.PostgresDB).
		WithRedisClient(app.RedisClient).
		WithAsynqClient(app.AsynqClient).
		WithAgentRepository(agentRepository).
		WithCloudStorageClient(storageThirdParty).
		WithUploadFileRepository(uploadfileRepository)

	// init gin handlers
	healthCheckHandler := healthcheck.New()
	uploadFileHandler := uploadFileHndlr.New().
		WithUploadFileService(uploadFileService)
	adminHandler := adminHndlr.New().
		WithAdminService(adminService)
	authHandler := auth.New().
		WithAdminService(adminService).
		WithAgentService(agentService)
	agentHandler := agentHndlr.New().
		WithAgentService(agentService)

	//init jwt middleware
	adminAuthMiddleware, err := middleware.NewAuthBuilder().
		WithAuthHandler(authHandler).
		WithRealm(constant.AdminUserType).
		Build()

	util.ContinueOrFatal(err)

	agentAuthMiddleware, err := middleware.NewAuthBuilder().
		WithAuthHandler(authHandler).
		WithRealm(constant.AgentUserType).
		Build()

	rbacMiddleware := middleware.NewRBACMiddleware().
		WithRoleToAccessRepository(roleToAccessRepository).
		WithAdminRepository(adminRepository)

	claimContextMiddleware := middleware.NewAccountSetupMiddleware().
		WithAdminRepository(adminRepository).
		WithAgentRepository(agentRepository)

	rateLimitMiddleware := middleware.NewRateLimitMiddleware().
		WithRedisDatabase(app.RedisClient)

	util.ContinueOrFatal(err)

	// init routes
	app.HTPPGinServer.Handler = gin.InitRoutes(
		healthCheckHandler,
		uploadFileHandler,
		adminHandler,
		adminAuthMiddleware,
		claimContextMiddleware,
		rbacMiddleware,
		agentHandler,
		agentAuthMiddleware,
		rateLimitMiddleware,
	)

	go func() {
		err = app.HTPPGinServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			util.ContinueOrFatal(err)
		}
	}()

	wait := util.GracefulShutdown(context.Background(), config.Env.App.GracefulShutdownTimeOut, map[string]util.Operation{
		"postgres connection": func(ctx context.Context) error {
			return pgDB.Close()
		},
		"redis connection": func(ctx context.Context) error {
			return app.RedisClient.Close()
		},
		"http gin server": func(ctx context.Context) error {
			return app.HTPPGinServer.Close()
		},
		"asynq client": func(ctx context.Context) error {
			return app.AsynqClient.Close()
		},
	})
	<-wait

}
