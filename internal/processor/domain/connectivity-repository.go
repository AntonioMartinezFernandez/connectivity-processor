package domain

type ConnectivityRepository interface {
	IsConnected(clientId string) (bool, error)
	GetConnectionId(clientId string) (string, error)
	SaveConnection(clientId string, connectionId string) error
	RemoveConnection(clientId string) error
}
