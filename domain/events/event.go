package events

import (
	"github.com/yakiza/BankZ/domain/events/types"
	"time"
)

// Event Entity
type Event struct {
	id        string
	version   int
	EventType types.Event
	timestamp time.Time
}

func (e Event) ID() string {
	return e.id
}

func (e Event) Version() int {
	return e.version
}

func (e Event) Type() types.Event {
	return e.EventType
}

func NewEvent(id string, version int, eventType types.Event, timestamp time.Time) Event {
	return Event{
		id:        id,
		version:   version,
		EventType: eventType,
		timestamp: timestamp,
	}
}
