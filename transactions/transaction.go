package transactions

import (
	"github.com/yakiza/BankZ/accounts"
	"time"
)

type Transaction struct {
	Type   TransactionType
	Status TransactionStatus

	FromAccount accounts.Account
	ToAccount   accounts.Account

	CreatedAt time.Time
}
