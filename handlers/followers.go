package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurmatthew/spotify-api-tool-go/services"
)

func FollowersHandler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	accessToken := req.Header.Get("access-token")
	clientToken := req.Header.Get("client-token")

	followers, err := services.GetFollowers(username, accessToken, clientToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching followers: %v", err), http.StatusInternalServerError)
		return
	}
	log.Println("successfully fetched followers info")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding followers: %v", err), http.StatusInternalServerError)
		return
	}
}
