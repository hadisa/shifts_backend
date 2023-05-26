package util

import (
	"time"

	"github.com/getsentry/sentry-go"
)

func SentryLogError(err error) {
	sentry.CaptureException(err)
	defer sentry.Flush(2 * time.Second)
}
