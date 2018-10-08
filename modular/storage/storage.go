package storage

import (
	"github.com/katzien/go-structure-examples/modular/beers"
	"github.com/katzien/go-structure-examples/modular/reviews"
)

// Type defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

// Storage defines the functionality of a data store for the beer service.
type Storage interface {
	SaveBeer(...beers.Beer) error
	SaveReview(...reviews.Review) error
	FindBeer(beers.Beer) ([]*beers.Beer, error)
	FindReview(reviews.Review) ([]*reviews.Review, error)
	FindBeers() []beers.Beer
	FindReviews() []reviews.Review
}

// DB is the "global" storage instance
var DB Storage

func NewStorage(t Type) error {
	var err error

	switch t {
	case Memory:
		DB = new(MemoryStorage)

	case JSON:
		// for the moment storage location for JSON files is the current working directory
		DB, err = NewJSONStorage("./data/")
		if err != nil {
			return err
		}
	}

	return nil
}
