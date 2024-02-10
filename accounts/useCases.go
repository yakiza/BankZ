package accounts

import (
	"context"
	"github.com/yakiza/BankZ/store"
	"log"
)

type AccountUseCases interface {
	// Create responsible for calling the respective persistence functions to create a new account.
	Create(ctx context.Context, account Account) error
}

var _ AccountUseCases = &UseCase{}

type UseCase struct {
	repo store.AccountRepository
}

func (uc *UseCase) Create(ctx context.Context, account Account) error {
	log.Printf("Creating Account for %s", account.Holder)

	err := uc.repo.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func NewAccountUseCase(repo store.AccountRepository) UseCase {
	return UseCase{
		repo: repo,
	}
}
