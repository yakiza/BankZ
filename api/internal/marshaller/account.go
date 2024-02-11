package marshaller

import (
	"github.com/yakiza/BankZ/domain/accounts"
)

type AccountDTO struct {
	Type   string `json:"type"`
	Holder string `json:"holder"`
}

func UnmarshalAccount(account AccountDTO) (accounts.Account, error) {
	var accType accounts.AccountType
	err := accType.FromString(account.Type)
	if err != nil {
		return accounts.Account{}, err
	}

	return accounts.Account{
		Type:   accType,
		Holder: account.Holder,
	}, err
}

func MarshalAccount(account accounts.Account) AccountDTO {
	return AccountDTO{
		Type:   account.Type.String(),
		Holder: account.Holder,
	}
}
