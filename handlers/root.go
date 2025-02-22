package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "spotify-api-tool-go /auth /profile /followers")
}
