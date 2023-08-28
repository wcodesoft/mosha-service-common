package tracing

import (
	"github.com/getsentry/sentry-go"
)

func SetupSentry(dsn string, release string, serverName string, sampleRate float64) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:           dsn,
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: sampleRate,
		Release:          release,
		ServerName:       serverName,
	})
	return err
}
