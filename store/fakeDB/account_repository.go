package fakeDB

import (
	"context"
	"errors"
	"github.com/yakiza/BankZ/accounts"
	"github.com/yakiza/BankZ/store"
	"sync"
)

var _ store.AccountRepository = &AccountRepository{}

type AccountRepository struct {
	mx       sync.Mutex
	accounts map[string]accounts.Account
}

func (rp *AccountRepository) CreateAccount(_ context.Context, account accounts.Account) error {
	rp.mx.Lock()
	defer rp.mx.Unlock()

	if _, ok := rp.accounts[account.Number]; !ok {
		rp.accounts[account.Number] = account
	} else {
		return errors.New("already exists")
	}

	return nil
}

func NewAccountRepository() AccountRepository {
	return AccountRepository{}
}
