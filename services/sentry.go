package services

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://e0a31fe5ccc8184b4ce85d3d2b6233f9@o4507486726258688.ingest.de.sentry.io/4507486728749136",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}
