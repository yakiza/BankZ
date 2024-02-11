package events

import (
	"github.com/yakiza/BankZ/domain/events/types"
	"time"
)

type Event struct {
	ID        string
	Version   int
	Type      types.Event
	Timestamp time.Time
}
