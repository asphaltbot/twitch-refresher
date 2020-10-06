package twitch

import (
	"github.com/asphaltbot/twitch-refresher/model"
	"github.com/asphaltbot/twitch-refresher/util/logging"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func MakeSampleRequest(config model.Config) bool {
	logging.InfoLine("Making sample twitch API request to see if our credentials are still valid")

	req, _ := http.NewRequest("GET", "https://api.twitch.tv/helix/search/channels?query=connorwrightkappa", nil)

	req.Header.Set("client-id", config.Credentials.TwitchClientID)
	req.Header.Set("Authorization", "Bearer "+config.Credentials.TwitchAccessToken)

	resp, err := httpClient.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logging.InfoLine("Received non-OK response from Twitch. Credentials need refreshing")
		return false
	}

	logging.InfoLine("Received OK response from Twitch. Credentials still valid.")
	return true

}
