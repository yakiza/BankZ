package events

import "github.com/yakiza/BankZ/events/types"

type EventStoreUseCases interface {
	// Append responsible for adding new events to the event store
	Append(event Event) error
	// QueryByTransactionID retrieving all events for a specific id
	QueryByTransactionID(transactionID string) ([]Event, error)
	// QueryByType retrieves all events of the type provided
	QueryByType(eventType types.Event) ([]Event, error)
}