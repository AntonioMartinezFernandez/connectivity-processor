package domain

type DisconnectedUsecase interface {
	Execute(event DisconnectedEvent) error
}
