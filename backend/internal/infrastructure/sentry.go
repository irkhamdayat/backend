package infrastructure

import (
	"fmt"
	"time"

	"github.com/Halalins/backend/config"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

func InitializeSentry(serviceName, serviceVersion string) {
	if config.Env.Env != "production" && config.Env.Env != "staging" {
		return
	}

	var (
		release = fmt.Sprintf("%s:%s", serviceName, serviceVersion)
	)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Env.Sentry.DSN,
		EnableTracing:    true,
		Release:          release,
		Environment:      config.Env.Env,
		TracesSampleRate: config.Env.Sentry.SampleRate,
		IgnoreErrors:     config.Env.Sentry.IgnoreErrors,
	})
	if err != nil {
		logrus.Fatalf("sentry.Init: %v", err)
	}

	hook, err := logrus_sentry.NewSentryHook(config.Env.Sentry.DSN, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		logrus.Fatalf("log_sentry.NewSentryHook: %v", err)
	}
	hook.StacktraceConfiguration.Enable = true
	hook.Timeout = config.Env.Sentry.Timeout
	hook.SetRelease(release)
	hook.SetEnvironment(config.Env.Env)
	err = hook.SetIgnoreErrors(config.Env.Sentry.IgnoreErrors...)
	if err != nil {
		logrus.Fatalf("log_sentry.SetIgnoreErrors: %v", err)
	}
	err = hook.SetSampleRate(float32(config.Env.Sentry.SampleRate))
	if err != nil {
		logrus.Fatalf("log_sentry.SetSampleRate: %v", err)
	}

	logrus.AddHook(hook)

}

func FlushSentry() {
	sentry.Flush(2 * time.Second)
}
