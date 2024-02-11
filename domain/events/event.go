package events

import (
	"github.com/yakiza/BankZ/domain/events/types"
	"time"
)

type Event struct {
	ID        string
	Type      types.Event
	Timestamp time.Time
}
