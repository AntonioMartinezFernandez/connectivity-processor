package application

import (
	http_domain "connectivity-processor/internal/http/domain"
	processor_domain "connectivity-processor/internal/processor/domain"
	"fmt"
)

func ProcessConnectedEvent(connectedRequest http_domain.ConnectedRequest, ce processor_domain.ConnectedUsecase) http_domain.DataResponse {
	fmt.Println("connected event received: ", connectedRequest)

	err := ce.Execute(processor_domain.ConnectedEvent(connectedRequest))

	res := http_domain.DataResponse{Data: err, Message: "Connected Event Processed"}
	return res
}
