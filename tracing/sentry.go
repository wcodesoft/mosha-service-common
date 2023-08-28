package tracing

import (
	"github.com/getsentry/sentry-go"
)

func SetupSentry(dsn string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	return err
}
