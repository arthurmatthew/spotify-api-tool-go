package router

import (
	"net/http"

	"github.com/arthurmatthew/spotify-api-tool-go/handlers"
	"github.com/arthurmatthew/spotify-api-tool-go/middleware"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/auth", handlers.AuthHandler)
	mux.HandleFunc("/profile", handlers.ProfileHandler)
	mux.HandleFunc("/followers", handlers.FollowersHandler)

	return middleware.RequestLogger(mux)
}
