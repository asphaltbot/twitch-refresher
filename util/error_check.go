package util

import (
	"github.com/asphaltbot/twitch-refresher/model"
	"github.com/asphaltbot/twitch-refresher/util/logging"
	"github.com/getsentry/sentry-go"
)

func CheckError(err error, config model.Config) {
	if err != nil {
		if config.Credentials.Sentry != "" {
			sentry.CaptureException(err)
			logging.InfoLine("Exception reported to Sentry")
		} else {
			logging.ErrorLine("An error occurred:", err)
		}
	}
}
