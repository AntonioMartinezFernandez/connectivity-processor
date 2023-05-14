package application

import (
	http_domain "connectivity-processor/internal/http/domain"
)

func HealthCheck() http_domain.HealthcheckResponse {
	response := http_domain.HealthcheckResponse{Status: "OK"}
	return response
}
