package events

import (
	"github.com/yakiza/BankZ/events/types"
	"time"
)

type Event struct {
	ID        string
	Type      types.Event
	Timestamp time.Time
}
