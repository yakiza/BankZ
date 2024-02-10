package accounts

import "fmt"

type AccountType int

const (
	main AccountType = iota
	savings
)

func (t *AccountType) String() string {
	var encoded string
	switch *t {
	case main:
		encoded = "main"
	case savings:
		encoded = "savings"
	default:
		return "Account type unrecognized"
	}

	return encoded
}

func (t *AccountType) FromString(accType string) error {
	switch accType {
	case "main":
		*t = main
	case "savings":
		*t = savings
	default:
		return fmt.Errorf("unrecognized account type")
	}

	return nil
}
