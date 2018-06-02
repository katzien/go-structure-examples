package database

import "github.com/katzien/structure-examples/modular/beers"
import "github.com/katzien/structure-examples/modular/reviews"

// StorageType defines available storage types
type StorageType int

const (
	// JSON will store data in JSON files saved on disk
	JSON StorageType = iota
	// Memory will store data in memory
	Memory
)

// DB is an interface to interact with data on multiple layered of data storage
var DB Storage

// Storage represents all possible actions available to deal with data
type Storage interface {
	SaveBeer(...beers.Beer) error
	SaveReview(...reviews.Review) error
	FindBeer(beers.Beer) ([]*beers.Beer, error)
	FindReview(reviews.Review) ([]*reviews.Review, error)
	FindBeers() []beers.Beer
	FindReviews() []reviews.Review
}

func NewStorage(storageType StorageType) error {
	var err error

	switch storageType {
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
