package main

import (
	"github.com/asphaltbot/twitch-refresher/api/twitch"
	"github.com/asphaltbot/twitch-refresher/api/twitchtokengenerator"
	"github.com/asphaltbot/twitch-refresher/model"
	"github.com/asphaltbot/twitch-refresher/util/logging"
	"github.com/getsentry/sentry-go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := model.ReadConfig()

	if config.Credentials.Sentry != "" {
		// initialise sentry
		err := sentry.Init(sentry.ClientOptions{
			Dsn: config.Credentials.Sentry,
		})

		if err != nil {
			logging.ErrorLine("Unable to initialise Sentry:", err)
		} else {
			logging.InfoLine("Initialised Sentry")
		}

	}

	if config.Credentials.TwitchClientID == "" {
		logging.FatalLine("You need to specify a client ID. You can get it from twitchtokengenerator.com")
	}

	if config.Credentials.TwitchAccessToken == "" {
		logging.FatalLine("You need to specify an access token to start with! You can use twitchtokengenerator.com")
	}

	if config.Credentials.TwitchRefreshToken == "" {
		logging.FatalLine("You need to specify a refresh token to start with! You can use twitchtokengenerator.com")
	}

	logging.InfoLine("Starting background threads")

	go func() {
		for {
			if !twitch.MakeSampleRequest(config) {
				twitchtokengenerator.GenerateNewCredentials(config)

				// todo: update all databases

			}

			time.Sleep(10 * time.Minute)
		}
	}()

	// prevent program from closing until we receive an explicit signal to do so
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc
	logging.InfoLine("Closing.")
	os.Exit(0)

}
