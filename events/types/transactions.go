package types

type Event int

const (
	EventTypeTransactionInitiated Event = iota
	EventTypeDeposit
	EventTypeWithdrawal
	EventTypeTransactionCompleted
	EventTypeTransactionCancelled
)
