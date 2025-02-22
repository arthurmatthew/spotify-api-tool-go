package middleware

import (
	"log"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s %s from %s", req.Method, req.URL.Path, req.Proto, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}
