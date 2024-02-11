package marshaller

import (
	accounts2 "github.com/yakiza/BankZ/domain/accounts"
)

type Account struct {
	Type   string `json:"type"`
	Holder string `json:"holder"`
}

func UnmarshalAccount(account Account) (accounts2.Account, error) {
	var accType accounts2.AccountType
	err := accType.FromString(account.Type)
	if err != nil {
		return accounts2.Account{}, err
	}

	return accounts2.Account{
		Type:   accType,
		Holder: account.Holder,
	}, err
}

func MarshalAccount(account accounts2.Account) Account {
	return Account{
		Type:   account.Type.String(),
		Holder: account.Holder,
	}
}
