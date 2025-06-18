package middleware

import (
	"log"
	"net/http"
	"time"
)

func Timing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("%s %s completed in %s", r.Method, r.URL.Path, elapsed)
	})
}
