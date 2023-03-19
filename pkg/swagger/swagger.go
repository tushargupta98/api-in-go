package swagger

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tushargupta98/api-in-go/config"
)

// SetupSwagger sets up the Swagger documentation for the HTTP server.
func SetupSwagger(r chi.Router, cfg config.SwaggerConfig) {
	// Serve the Swagger UI at the /swagger/ path.
	r.Handle(cfg.SwaggerUiPath+"*", httpSwagger.Handler(
		httpSwagger.URL(cfg.JsonPath), //The url pointing to API definition
	))

	// Rendering Swagger Docs.
	r.Handle(cfg.DocsPath+"*", http.StripPrefix(cfg.DocsPath, http.FileServer(http.Dir(cfg.StaticDir))))
}
