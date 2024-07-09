package bootstrap

import (
	//Prebuild Library
	"context"
	"os"

	//External Library
	"github.com/go-gomail/gomail"
	"github.com/hibiken/asynq"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	//Internal Library: Common
	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/infrastructure"

	//Internal Library: Thirdparty
	"github.com/Halalins/backend/internal/thirdparty/mailer"
	"github.com/Halalins/backend/internal/thirdparty/onesignal"

	//Internal Library: Repository
	"github.com/Halalins/backend/internal/repository/notification"

	//Internal Library: Service
	mailerSvc "github.com/Halalins/backend/internal/service/mailer"
	notificationSvc "github.com/Halalins/backend/internal/service/notification"

	//Internal Library: Handler
	asynqHndlr "github.com/Halalins/backend/internal/handler/asynq"
	mailerHndlr "github.com/Halalins/backend/internal/handler/asynq/mailer"
	notificationHndlr "github.com/Halalins/backend/internal/handler/asynq/notification"
)

type AppWorker struct {
	PostgresDB     *gorm.DB
	RedisClient    *redis.Client
	AysnqServer    *asynq.Server
	AsynqScheduler *asynq.Scheduler
	GomailDialer   *gomail.Dialer
	AsynqClient    *asynq.Client
	I18nBundle     *i18n.Bundle
}

func greetingWorker() {
	logrus.Info(config.Env.App.Name, " Worker Started with PID: ", os.Getpid())
}

func StartWorker() {
	greetingWorker()

	infrastructure.InitializeSentry(ServiceName+"-worker", ServiceVersion)

	app := AppWorker{
		PostgresDB:     infrastructure.InitializePostgresConn(),
		RedisClient:    infrastructure.InitializeRedisConn(),
		AysnqServer:    infrastructure.InitializeAsynqServer(),
		AsynqScheduler: infrastructure.InitializeAsynqScheduler(),
		AsynqClient:    infrastructure.InitializeAsynqClient(),
		GomailDialer:   infrastructure.InitializeGomailDialer(),
		I18nBundle:     infrastructure.InitializeI18nBundle(),
	}

	_ = errmapper.Initialize().WithI18nBundle(app.I18nBundle, "notifications")

	pgDB, err := app.PostgresDB.DB()
	util.ContinueOrFatal(err)

	// init third party
	mailerThirdParty := mailer.New().
		WithGomailDialer(app.GomailDialer)
	onesignalThirdParty := onesignal.New().
		WithI18nBundle(app.I18nBundle)

	// init repositories
	notificationHistoryRepository := notification.New().
		WithPostgresDB(app.PostgresDB)

	// init services
	notificationService := notificationSvc.New().
		WithPostgresDB(app.PostgresDB).
		WithI18nBundle(app.I18nBundle).
		WithOnesignalRequester(onesignalThirdParty).
		WithNotificationHistoryRepository(notificationHistoryRepository)

	mailerService := mailerSvc.New().
		WithMailerThirdParty(mailerThirdParty)

	// init asynq handlers
	mailerHandler := mailerHndlr.New().
		WithMailerService(mailerService)

	notificationHandler := notificationHndlr.New().
		WithNotificationService(notificationService)

	mux := asynqHndlr.NewServeMuxBuilder().
		WithMiddleware(asynqHndlr.LoggingMiddleware()).
		WithMailerHandler(mailerHandler).
		WithNotificationService(notificationHandler).
		Build()

	go func() {
		asynqHndlr.RegisterSchedulerTask(app.AsynqScheduler)
		if err = app.AsynqScheduler.Run(); err != nil {
			logrus.Fatal(err)
		}
	}()

	go func() {
		if err = app.AysnqServer.Run(mux); err != nil {
			logrus.Fatal(err)
		}
	}()

	wait := util.GracefulShutdown(context.Background(), config.Env.App.GracefulShutdownTimeOut, map[string]util.Operation{
		"postgressql connection": func(ctx context.Context) error {
			return pgDB.Close()
		},
		"asynq connection": func(ctx context.Context) error {
			app.AysnqServer.Stop()
			return nil
		},
		"asynq client": func(ctx context.Context) error {
			return app.AsynqClient.Close()
		},
		"asynq scheduler connection": func(ctx context.Context) error {
			app.AsynqScheduler.Shutdown()
			return nil
		},
	})
	<-wait
}
