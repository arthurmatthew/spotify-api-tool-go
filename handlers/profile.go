package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurmatthew/spotify-api-tool-go/services"
)

func ProfileHandler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	accessToken := req.Header.Get("access-token")
	clientToken := req.Header.Get("client-token")

	profile, err := services.GetProfile(username, accessToken, clientToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching profile: %v", err), http.StatusInternalServerError)
		return
	}
	log.Println("successfully fetched profile info")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding profile: %v", err), http.StatusInternalServerError)
		return
	}
}
