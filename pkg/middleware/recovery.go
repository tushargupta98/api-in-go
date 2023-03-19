package middleware

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tushargupta98/api-in-go/logger"
)

// AddRecovery adds recovery middleware to the router.
func AddRecovery(router *chi.Mux) {
	router.Use(recovery)
}

// recovery middleware function.
func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// Log the panic with the error message.
				logger.Logger.WithField("error", rec).Error("recovered from panic")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
