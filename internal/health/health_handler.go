package health

import (
	"net/http"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// List godoc
// @Summary Health check endpoint
// @Description Returns the status of the service.
// @Tags API Health Check
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheckEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
