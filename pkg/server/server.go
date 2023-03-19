package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/tushargupta98/api-in-go/cache"
	"github.com/tushargupta98/api-in-go/config"
	"github.com/tushargupta98/api-in-go/internal/domain"
	"github.com/tushargupta98/api-in-go/internal/health"
	"github.com/tushargupta98/api-in-go/logger"
	config_middleware "github.com/tushargupta98/api-in-go/pkg/middleware"
	"github.com/tushargupta98/api-in-go/pkg/swagger"
)

// Server is the main HTTP server struct.
type Server struct {
	Router *chi.Mux
}

// NewServer returns a new instance of the HTTP server.
func NewServer() *Server {
	r := chi.NewRouter()

	return &Server{
		Router: r,
	}
}

// SetupMiddleware sets up the middleware for the server.
func (s *Server) SetupMiddleware() {
	// Add CORS middleware.
	config_middleware.AddCORS(s.Router)

	// Add logging middleware.
	config_middleware.AddLogging(s.Router)

	// Add recovery middleware.
	config_middleware.AddRecovery(s.Router)

	// Add authenticator middleware.
	authenticator := config_middleware.ApiKeyAuthenticator("unique-api-key")

	config_middleware.AddAuthentication(s.Router, authenticator)

}

// SetupRoutes sets up the routes for the server.
func (s *Server) SetupRoutes(db *sqlx.DB, cache cache.RedisClient, config config.Config) {
	// Add middleware to set up the database connection
	s.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			type dbKeyType string
			const dbKey dbKeyType = "db"

			ctx = context.WithValue(ctx, dbKey, db)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	// Add health check endpoint to the api
	s.Router.Get(config.SwaggerConfig.BasePath+"/health", health.HealthCheckEndpoint)

	// Add swagger endpoints to the api
	swagger.SetupSwagger(s.Router, config.SwaggerConfig)

	// Add entity endpoints to the api
	s.Router.Route(config.SwaggerConfig.BasePath, func(s chi.Router) {
		domain.DomainRouter(s, db, cache)
	})
}

// ListenAndServe starts the server and listens for incoming HTTP requests.
func (s *Server) ListenAndServe(addr string) error {
	logger.Logger.Infof("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.Router)
}
