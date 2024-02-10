package transactions

import (
	"context"
	"fmt"
)

type TransactionUseCases interface {
	Withdraw(ctx context.Context, amount int, accountNumber string) error
	Deposit(ctx context.Context, amount int, accountNumber string) error
}

var _ TransactionUseCases = &UseCases{}

type UseCases struct {
}

func (c UseCases) Withdraw(ctx context.Context, amount int, accountNumber string) error {
	if amount <= 0 {
		return fmt.Errorf("incorrect amount")
	}

	return nil
}

func (c UseCases) Deposit(ctx context.Context, amount int, accountNumber string) error {
	if amount <= 0 {
		return fmt.Errorf("incorrect amount")
	}

	return nil
}

func NewTransactionUseCases() UseCases {
	return UseCases{}
}
