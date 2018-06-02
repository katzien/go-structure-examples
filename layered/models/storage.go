package models

import "github.com/katzien/structure-examples/layered/storage"

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
	SaveBeer(...Beer) error
	SaveReview(...Review) error
	FindBeer(Beer) ([]*Beer, error)
	FindReview(Review) ([]*Review, error)
	FindBeers() []Beer
	FindReviews() []Review
}

func NewStorage(storageType StorageType) error {
	var err error

	switch storageType {
	case Memory:
		DB = new(storage.Memory)

	case JSON:
		// for the moment storage location for JSON files is the current working directory
		DB, err = storage.NewJSON("./data/")
		if err != nil {
			return err
		}
	}

	return nil
}
