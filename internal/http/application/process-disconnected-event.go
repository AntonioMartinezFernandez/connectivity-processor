package application

import (
	http_domain "connectivity-processor/internal/http/domain"
	processor_domain "connectivity-processor/internal/processor/domain"
	"fmt"
)

func ProcessDisconnectedEvent(disconnectedRequest http_domain.DisconnectedRequest, de processor_domain.DisconnectedUsecase) http_domain.DataResponse {
	fmt.Println("disconnected event received: ", disconnectedRequest)

	err := de.Execute(processor_domain.DisconnectedEvent(disconnectedRequest))

	res := http_domain.DataResponse{Data: err, Message: "Disconnected Event Processed"}
	return res
}
