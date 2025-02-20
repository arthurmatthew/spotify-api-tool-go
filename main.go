package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurmatthew/spotify-api-tool-go/services"
)

const PORT int = 8765

type TokensResponse struct {
	ClientToken string `json:"client_token"`
	AccessToken string `json:"access_token"`
}

func auth(w http.ResponseWriter, req *http.Request) {
	clientTokenChan := make(chan *services.ClientTokenObject)
	accessTokenChan := make(chan *services.AccessTokenObject)
	errorChan := make(chan error)

	go func() {
		accessToken, err := services.GetAccessTokenObject()
		if err != nil {
			errorChan <- fmt.Errorf("error fetching access token: %v", err)
			return
		}
		accessTokenChan <- accessToken
	}()

	accessToken := <-accessTokenChan
	if accessToken == nil {
		http.Error(w, "failed to fetch access token", http.StatusInternalServerError)
		return
	}
	log.Printf("successfully fetched access token")

	go func() {
		clientToken, err := services.GetClientTokenObject(*accessToken)
		if err != nil {
			errorChan <- fmt.Errorf("error fetching client token: %v", err)
			return
		}
		clientTokenChan <- clientToken
	}()

	clientToken := <-clientTokenChan
	if clientToken == nil {
		http.Error(w, "failed to fetch client token", http.StatusInternalServerError)
		return
	}
	log.Printf("successfully fetched client token")

	response := TokensResponse{
		ClientToken: clientToken.GrantedToken.Token,
		AccessToken: accessToken.AccessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

// func user(w http.ResponseWriter, req *http.Request) {

// }

// func followers(w http.ResponseWriter, req *http.Request) {

// }

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s from %s", r.Method, r.URL.Path, r.Proto, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", auth)
	// mux.HandleFunc("/user", user)
	// mux.HandleFunc("/followers", followers)

	loggedMux := requestLogger(mux)

	addr := fmt.Sprintf(":%d", PORT)

	log.Println("Server listening on", addr)
	err := http.ListenAndServe(addr, loggedMux)
	if err != nil {
		log.Fatal(err)
	}
}
