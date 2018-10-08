package reviews

import (
	"errors"
	"time"
)

// Review defines a beer review
type Review struct {
	ID        string    `json:"id"`
	BeerID    int       `json:"beer_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given beer ID.
	GetAll(int) []Review
	// Add saves a given review.
	Add(Review) error
}
