package memory

import (
	"time"
)

// Review defines a beer review
type Review struct {
	ID        string
	BeerID    int
	FirstName string
	LastName  string
	Score     int
	Text      string
	Created   time.Time
}
