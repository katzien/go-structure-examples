package reviews

import (
	"time"
	"errors"
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

// Repository provides access to the reviews.
type Repository interface {
	GetAll(beerID int) []Review
	Add(review Review) error
}
