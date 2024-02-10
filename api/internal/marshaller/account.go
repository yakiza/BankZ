package marshaller

import (
	"github.com/yakiza/BankZ/accounts"
)

type Account struct {
	Type   string   `json:"type"`
	Holder Customer `json:"holder"`
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func UnmarshalAccount(account Account) accounts.Account {
	return accounts.Account{
		Type:   account.Type,
		Holder: account.Holder,
	}
}

func MarshalAccount(account accounts.Account) Account {
	return Account{
		Type:   account.Type,
		Holder: account.Holder,
	}
}
