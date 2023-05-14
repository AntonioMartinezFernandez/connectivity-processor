package infrastructure

import (
	processor_domain "connectivity-processor/internal/processor/domain"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter(ce processor_domain.ConnectedUsecase, de processor_domain.DisconnectedUsecase) *mux.Router {
	gorillaMuxRouter := mux.NewRouter()

	// Healthcheck
	gorillaMuxRouter.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		HealthCheckHandler(w, r)
	}).Methods("GET")

	// Connectivity
	gorillaMuxRouter.HandleFunc("/connected", func(w http.ResponseWriter, r *http.Request) {
		ConnectHandler(w, r, ce)
	}).Methods("POST")
	gorillaMuxRouter.HandleFunc("/disconnected", func(w http.ResponseWriter, r *http.Request) {
		DisconnectHandler(w, r, de)
	}).Methods("POST")

	return gorillaMuxRouter
}
