package accounts

import (
	"github.com/yakiza/BankZ/domain/command"
	"github.com/yakiza/BankZ/domain/events"
	"github.com/yakiza/BankZ/store"
	"github.com/yakiza/BankZ/store/eventstore"
)

// AccountUseCases describes the Account UseCases
type AccountUseCases interface {
	CommandHandler(cmd command.Command) ([]events.Event, error)
}

var _ AccountUseCases = &UseCase{}

type UseCase struct {
	eventStore eventstore.EventStore
	repository store.AccountRepository
}

func (uc UseCase) CommandHandler(cmd command.Command) ([]events.Event, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountUseCase(eventStore eventstore.EventStore, repository store.AccountRepository) UseCase {
	return UseCase{
		eventStore: eventStore,
		repository: repository,
	}
}
