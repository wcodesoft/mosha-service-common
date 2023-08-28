package tracing

import (
	"github.com/getsentry/sentry-go"
)

func SetupSentry(dsn string, release string, serverName string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		EnableTracing:    true,
		TracesSampleRate: 0.2,
		Release:          release,
		ServerName:       serverName,
	})
	return err
}
