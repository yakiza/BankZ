package transactions

import "fmt"

type TransactionType int

const (
	deposit TransactionType = iota
	withdrawal
	transfer
)

func (t *TransactionType) String() string {
	var encoded string
	switch *t {
	case deposit:
		encoded = "deposit"
	case withdrawal:
		encoded = "withdrawal"
	case transfer:
		encoded = "transfer"
	default:
		return "Transaction type unrecognized"
	}

	return encoded
}

func (t *TransactionType) FromString(transType string) error {
	switch transType {
	case "deposit":
		*t = deposit
	case "withdrawal":
		*t = withdrawal
	case "transfer":
		*t = transfer
	default:
		return fmt.Errorf("unrecognized transaction type")
	}

	return nil
}
