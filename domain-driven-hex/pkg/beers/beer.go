package beers

import (
	"time"
	"errors"
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

// Repository provides access to the list of beers.
type Repository interface {
	GetAll() []Beer
	Get(id int) (Beer, error)
	Add(Beer) error
}