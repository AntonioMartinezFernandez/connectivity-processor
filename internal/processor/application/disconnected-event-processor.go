package application

import (
	processor_domain "connectivity-processor/internal/processor/domain"
	"encoding/json"
	"errors"
	"fmt"
)

type DisconnectedEventProcessor struct {
	connRepo processor_domain.ConnectivityRepository
}

func NewDisconnectedEventProcessor(connRepo processor_domain.ConnectivityRepository) *DisconnectedEventProcessor {
	return &DisconnectedEventProcessor{
		connRepo: connRepo,
	}
}

func (dep *DisconnectedEventProcessor) Execute(disconnectedEvent processor_domain.DisconnectedEvent) error {
	eventAsJsonBytesArray, err := json.Marshal(disconnectedEvent)
	if err != nil {
		return errors.New("error marshalling disconnected event")
	}

	connectionId, getConnectionIdError := dep.connRepo.GetConnectionId(disconnectedEvent.ClientID)
	if getConnectionIdError != nil {
		return errors.New("error retrieving connectionId from connectivity repository")
	}

	if connectionId == disconnectedEvent.SessionIdentifier {
		fmt.Println("### PUBLISH THING DISCONNECTED EVENT", string(eventAsJsonBytesArray))
		removeConnectionError := dep.connRepo.RemoveConnection(disconnectedEvent.ClientID)
		if removeConnectionError != nil {
			return errors.New("error removing connection from connectivity repository")
		}
	}

	return nil
}
