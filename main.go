package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurmatthew/spotify-api-tool-go/router"
)

func main() {
	addr := fmt.Sprintf(":%d", 8765)
	log.Println("Server listening on", addr)

	err := http.ListenAndServe(addr, router.SetupRouter())
	if err != nil {
		log.Fatal(err)
	}
}
