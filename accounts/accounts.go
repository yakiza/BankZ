package accounts

import (
	"context"
	"github.com/yakiza/BankZ/store"
	"log"
	"time"
)

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

func NewUseCase(repo store.AccountRepository) UseCase {
	return UseCase{
		repo: repo,
	}
}

type Account struct {
	Number    string
	Type      string
	Holder    string
	CreatedAt time.Time
	Balance   string
}
