package memory

import (
	"time"
)

// Review defines the storage form of a beer review
type Review struct {
	ID        string
	BeerID    string
	FirstName string
	LastName  string
	Score     int
	Text      string
	Created   time.Time
}
