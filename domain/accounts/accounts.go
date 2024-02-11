package accounts

import (
	"time"
)

// Account Entity/Aggregate
type Account struct {
	Number    string
	Type      AccountType
	Holder    string
	CreatedAt time.Time
	Balance   string
}
