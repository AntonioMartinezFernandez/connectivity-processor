package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	http_application "connectivity-processor/internal/http/application"
	http_domain "connectivity-processor/internal/http/domain"
	processor_domain "connectivity-processor/internal/processor/domain"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request, ce processor_domain.ConnectedUsecase) {
	fmt.Println("processing 'connected' event...")

	// Read request body
	requestBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendResponseUnprocessableEntity(w)
	}

	// Unmarshall body
	parsedRequestBody := http_domain.ConnectedRequest{}
	if err := json.Unmarshal(requestBody, &parsedRequestBody); err != nil {
		fmt.Println("error unmarshalling request body: ", err)
		SendResponseUnprocessableEntity(w)
	} else {
		// Run usecase
		handlerResponse := http_application.ProcessConnectedEvent(parsedRequestBody, ce)
		if handlerResponse.Data != nil {
			fmt.Println("process connected handler error: ", handlerResponse)
			SendResponseInternalServerError(w)
		} else {
			SendResponseWithData(w, handlerResponse)
		}
	}
}
