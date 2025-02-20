package services

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type AccessTokenObject struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
	ClientId    string `json:"clientId"`
}

type ClientTokenObject struct {
	ResponseType string `json:"response_type"`
	GrantedToken struct {
		Token               string   `json:"token"`
		ExpiresAfterSeconds int      `json:"expires_after_seconds"`
		RefreshAfterSeconds int      `json:"refresh_after_seconds"`
		Domains             []Domain `json:"domains"`
	} `json:"granted_token"`
}

type Domain struct {
	Domain string `json:"domain"`
}

type ClientData struct {
	ClientVersion string `json:"client_version"`
	ClientID      string `json:"client_id"`
	JSSDKData     struct {
		DeviceBrand string `json:"device_brand"`
		DeviceModel string `json:"device_model"`
		OS          string `json:"os"`
		OSVersion   string `json:"os_version"`
		DeviceID    string `json:"device_id"`
		DeviceType  string `json:"device_type"`
	} `json:"js_sdk_data"`
}

type RequestBody struct {
	ClientData ClientData `json:"client_data"`
}

func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var result []byte = make([]byte, length)

	for i := 0; i < length; i++ {
		randomByte := make([]byte, 1)
		_, err := rand.Read(randomByte)
		if err != nil {
			return "", fmt.Errorf("error generating random byte: %v", err)
		}

		result[i] = charset[randomByte[0]%byte(len(charset))]
	}

	return string(result), nil
}

func GetAccessTokenObject() (*AccessTokenObject, error) {
	resp, err := http.Get("https://open.spotify.com/")
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %v", err)
	}

	tokenRegex := regexp.MustCompile(`<script id="session" data-testid="session" type="application\/json">(.+?)<\/script>`)
	matches := tokenRegex.FindStringSubmatch(string(body))

	var tokenData AccessTokenObject
	err = json.Unmarshal([]byte(matches[1]), &tokenData)
	if err != nil {
		return nil, fmt.Errorf("parse access token JSON error: %v", err)
	}

	return &tokenData, nil
}

func GetClientTokenObject(accessTokenObject AccessTokenObject) (*ClientTokenObject, error) {
	clientId := accessTokenObject.ClientId
	deviceId, err := generateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("device id generation error: %v", err)
	}

	requestBody := RequestBody{
		ClientData: ClientData{
			ClientVersion: "1.2.45.254.g74c8e3c5",
			ClientID:      clientId,
			JSSDKData: struct {
				DeviceBrand string `json:"device_brand"`
				DeviceModel string `json:"device_model"`
				OS          string `json:"os"`
				OSVersion   string `json:"os_version"`
				DeviceID    string `json:"device_id"`
				DeviceType  string `json:"device_type"`
			}{
				DeviceBrand: "unknown",
				DeviceModel: "unknown",
				OS:          "windows",
				OSVersion:   "NT 10.0",
				DeviceID:    deviceId,
				DeviceType:  "computer",
			},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "https://clienttoken.spotify.com/v1/clienttoken", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failure: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response non-OK status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var tokenResponse ClientTokenObject
	err = json.Unmarshal(bodyBytes, &tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response JSON: %v", err)
	}

	return &tokenResponse, nil
}
