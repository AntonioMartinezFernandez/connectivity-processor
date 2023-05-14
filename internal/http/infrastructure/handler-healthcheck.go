package infrastructure

import (
	"net/http"

	http_application "connectivity-processor/internal/http/application"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Get Healthcheck response
	handlerResponse := http_application.HealthCheck()

	// Response
	SendResponseHealthcheck(w, handlerResponse)
}
