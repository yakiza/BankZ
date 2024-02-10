package accounts

import "fmt"

type accountType int

const (
	main accountType = iota
	savings
)

func (t *accountType) ToString() string {
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

func (t *accountType) FromString(accType string) error {
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
