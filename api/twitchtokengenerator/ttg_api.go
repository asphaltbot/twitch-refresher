package twitchtokengenerator

import (
	"encoding/json"
	"github.com/asphaltbot/twitch-refresher/model"
	"github.com/asphaltbot/twitch-refresher/util"
	"github.com/asphaltbot/twitch-refresher/util/logging"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type GenerateResponse struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh"`
	ClientID     string `json:"client_id"`
}

func GenerateNewCredentials(config model.Config) {
	logging.InfoLine("Generating new credentials")

	req, _ := http.NewRequest("GET", "https://twitchtokengenerator.com/api/refresh/"+config.Credentials.TwitchRefreshToken, nil)
	resp, err := httpClient.Do(req)

	util.CheckError(err, config)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	util.CheckError(err, config)

	var generateResponse GenerateResponse
	err = json.Unmarshal(bodyBytes, &generateResponse)
	util.CheckError(err, config)

	config.Credentials.TwitchRefreshToken = generateResponse.RefreshToken
	config.Credentials.TwitchAccessToken = generateResponse.AccessToken
	config.Credentials.TwitchClientID = generateResponse.ClientID

	config.Save()

}
