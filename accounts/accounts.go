package accounts

import (
	"time"
)

type Account struct {
	Number    string
	Type      AccountType
	Holder    string
	CreatedAt time.Time
	Balance   string
}
