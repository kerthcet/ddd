package po

import (
	"time"
)

// Event po
type Event struct {
	ID            int
	CreatedAt     time.Time
	DomainEventID string
	Metadata      string
}
