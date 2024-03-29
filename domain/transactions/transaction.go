package transactions

import (
	"github.com/yakiza/BankZ/domain/accounts"
	"time"
)

type Transaction struct {
	Type        TransactionType
	Status      TransactionStatus
	Description string

	FromAccount accounts.Account
	ToAccount   accounts.Account

	Amount    string
	CreatedAt time.Time
}
