package beers

import (
	"errors"
	"time"
)

// Beer defines the properties of a reviewable beer
type Beer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}

// ErrUnknown is used when a beer could not be found.
var ErrUnknown = errors.New("unknown beer")
var ErrDuplicate = errors.New("beer already exists")

// Repository provides access to the beer storage.
type Repository interface {
	// GetAll returns all beers saved in storage.
	GetAll() []Beer
	// Get returns the beer with given ID.
	Get(int) (Beer, error)
	// Add saves a given beer to the repository.
	Add(Beer) error
}
