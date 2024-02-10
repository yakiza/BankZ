package accounts

import "context"

type AccountUseCases interface {
	// Create responsible for calling the respective persistence functions to create a new account.
	Create(ctx context.Context, account Account) error
}
