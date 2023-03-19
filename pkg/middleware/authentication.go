package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// Authenticator is a function that returns true if the request is authenticated, false otherwise.
type Authenticator func(r *http.Request) bool

type contextKey string

// RequireAuthentication returns middleware that requires an authenticated request.
func RequireAuthentication(authenticator Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "swagger-ui") || strings.Contains(r.URL.Path, "swagger-docs") {
				// Allow unauthenticated requests to the Swagger UI endpoint.
				next.ServeHTTP(w, r)
				return
			}

			if authenticator(r) {
				ctx := context.WithValue(r.Context(), contextKey("authenticated"), true)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			}
		})
	}
}

// APIKeyAuthenticator provider
func ApiKeyAuthenticator(apiKey string) Authenticator {
	return func(r *http.Request) bool {
		key := r.Header.Get("X-API-KEY")
		return key == apiKey
	}
}

// AddAuthentication adds authentication middleware to the router.
func AddAuthentication(router *chi.Mux, authenticator Authenticator) {
	router.Use(RequireAuthentication(authenticator))
}
