package model

import (
	"encoding/json"
	"github.com/asphaltbot/twitch-refresher/util/logging"
	"io/ioutil"
	"os"
)

type Config struct {
	Hosts []struct {
		IP            string `json:"ip"`
		Port          int    `json:"port"`
		Username      string `json:"username"`
		Password      string `json:"password"`
		Database      string `json:"db"`
		DatabaseTable string `json:"table"`
	} `json:"hosts"`
	Credentials struct {
		Sentry             string `json:"sentry"`
		TwitchAccessToken  string `json:"twitch_access_token"`
		TwitchClientID     string `json:"twitch_client_id"`
		TwitchRefreshToken string `json:"twitch_refresh_token"`
	}
}

func ReadConfig() Config {

	configFile, err := os.OpenFile("config.json", os.O_RDONLY, 0644)

	if err != nil {
		logging.FatalLine("Could not open config.json:", err)
	}

	defer configFile.Close()

	fileBytes, err := ioutil.ReadAll(configFile)

	if err != nil {
		logging.FatalLine(err)
	}

	var config Config
	err = json.Unmarshal(fileBytes, &config)

	if err != nil {
		logging.FatalLine(err)
	}

	return config

}

func (c Config) Save() {
	logging.InfoLine("Saving new config to config.json")
	configFile, err := os.OpenFile("config.json", os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		logging.FatalLine(err)
	}

	defer configFile.Close()

	structBytes, err := json.Marshal(c)

	if err != nil {
		logging.FatalLine(err)
	}

	_, err = configFile.Write(structBytes)

	if err != nil {
		logging.FatalLine(err)
	}

	logging.InfoLine("config.json file successfully saved.")

}
