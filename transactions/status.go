package transactions

import "fmt"

type TransactionStatus int

const (
	inProgress TransactionStatus = iota
	completed
	cancelled
	onHold
)

func (ts *TransactionStatus) String() string {
	var encoded string
	switch *ts {
	case inProgress:
		encoded = "inProgress"
	case completed:
		encoded = "completed"
	case cancelled:
		encoded = "cancelled"
	case onHold:
		encoded = "onHold"
	default:
		return "Transaction status unrecognized"
	}

	return encoded
}

func (ts *TransactionStatus) FromString(status string) error {
	switch status {
	case "inProgress":
		*ts = inProgress
	case "completed":
		*ts = completed
	case "cancelled":
		*ts = cancelled
	case "onHold":
		*ts = onHold
	default:
		return fmt.Errorf("unrecognized transaction status")
	}

	return nil
}
