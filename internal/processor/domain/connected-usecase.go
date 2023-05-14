package domain

type ConnectedUsecase interface {
	Execute(event ConnectedEvent) error
}
