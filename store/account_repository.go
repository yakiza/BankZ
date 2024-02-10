package store

import (
	"context"
	"github.com/yakiza/BankZ/accounts"
)

type AccountRepository interface {
	// CreateAccount persistence layer function that inserts an entry of a new account in to the respective DB
	CreateAccount(ctx context.Context, account accounts.Account) error
}
