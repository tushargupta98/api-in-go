package middleware

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tushargupta98/api-in-go/logger"
)

// AddLogging adds logging middleware to the router.
func AddLogging(router *chi.Mux) {
	router.Use(logging())
}

// logging is middleware that logs each request using the specified logger.
func logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Logger.Infof("Handling request %s %s", r.Method, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}
