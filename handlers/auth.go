package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurmatthew/spotify-api-tool-go/services"
)

type TokensResponse struct {
	ClientToken services.ClientTokenObject `json:"client_token"`
	AccessToken services.AccessTokenObject `json:"access_token"`
}

func AuthHandler(w http.ResponseWriter, req *http.Request) {
	accessToken, err := services.GetAccessTokenObject()
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching access token: %v", err), http.StatusInternalServerError)
		return
	}
	log.Println("successfully fetched access token")

	clientToken, err := services.GetClientTokenObject(*accessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching client token: %v", err), http.StatusInternalServerError)
		return
	}
	log.Println("successfully fetched client token")

	response := TokensResponse{
		ClientToken: *clientToken,
		AccessToken: *accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
