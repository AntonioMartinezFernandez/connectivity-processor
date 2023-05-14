package application

import (
	processor_domain "connectivity-processor/internal/processor/domain"
	"encoding/json"
	"errors"
	"fmt"
)

type ConnectedEventProcessor struct {
	connRepo processor_domain.ConnectivityRepository
}

func NewConnectedEventProcessor(connRepo processor_domain.ConnectivityRepository) *ConnectedEventProcessor {
	return &ConnectedEventProcessor{
		connRepo: connRepo,
	}
}

func (cep *ConnectedEventProcessor) Execute(connectedEvent processor_domain.ConnectedEvent) error {
	eventAsJsonBytesArray, marshalError := json.Marshal(connectedEvent)
	if marshalError != nil {
		return errors.New("error marshalling connected event")
	}

	thingIsConnected, getConnectedEventError := cep.connRepo.IsConnected(connectedEvent.ClientID)
	if getConnectedEventError != nil {
		return errors.New("error retrieving connection status from connectivity repository")
	}

	saveConnectionError := cep.connRepo.SaveConnection(connectedEvent.ClientID, connectedEvent.SessionIdentifier)
	if saveConnectionError != nil {
		fmt.Println("### PUBLISH THING CONNECTED EVENT", string(eventAsJsonBytesArray))
		return errors.New("error saving connection status with connectivity repository")
	}

	if !thingIsConnected {
		fmt.Println("### PUBLISH THING CONNECTED EVENT", string(eventAsJsonBytesArray))
	}

	return nil
}
