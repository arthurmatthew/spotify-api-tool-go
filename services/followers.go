package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/arthurmatthew/spotify-api-tool-go/models"
)

func GetFollowers(username string, accessToken string, clientToken string) ([]models.Follower, error) {
	url := fmt.Sprintf("https://spclient.wg.spotify.com/user-profile-view/v3/profile/%s/followers?market=US", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header = http.Header{
		"Priority":           {"u=1, i"},
		"Sec-Ch-Ua":          {`"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`},
		"Sec-Ch-Ua-mobile":   {"?0"},
		"Sec-Ch-Ua-Platform": {`"Windows"`},
		"Sec-Fetch-Dest":     {"empty"},
		"Sec-Fetch-Mode":     {"cors"},
		"Sec-Fetch-Site":     {"same-site"},
		"Referer":            {"https://open.spotify.com/"},
		"Referrer-Policy":    {"strict-origin-when-cross-origin"},
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("App-Platform", "WebPlayer")
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %v`, accessToken))
	req.Header.Set("Client-Token", clientToken)
	req.Header.Set("Spotify-App-Version", "1.2.45.254.g74c8e3c5")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failure: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response non-OK status: %v", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var followersResponse models.FollowersResponse
	err = json.Unmarshal(bodyBytes, &followersResponse)
	if err != nil {
		return nil, fmt.Errorf("error reading json: %v", err)
	}

	return followersResponse.Followers, nil
}
