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
		encoded = "in_progress"
	case completed:
		encoded = "completed"
	case cancelled:
		encoded = "cancelled"
	case onHold:
		encoded = "on_hold"
	default:
		return "Transaction status unrecognized"
	}

	return encoded
}

func (ts *TransactionStatus) FromString(status string) error {
	switch status {
	case "in_progress":
		*ts = inProgress
	case "completed":
		*ts = completed
	case "cancelled":
		*ts = cancelled
	case "on_hold":
		*ts = onHold
	default:
		return fmt.Errorf("unrecognized transaction status")
	}

	return nil
}
