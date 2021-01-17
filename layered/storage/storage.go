package storage

// This file belongs to storage package instead of in models.
// models only contain the structure that is universally used by all packages

import "github.com/katzien/go-structure-examples/layered/models"

// StorageType defines available storage types
// this will stutter. think of this package as a dal (data access layer)
type StorageType int

const (
	// JSONType will store data in JSON files saved on disk
	JSONType StorageType = iota
	// MemoryType will store data in memory
	MemoryType
)

// DB is an interface to interact with data on multiple layered of data storage
var DB Storage

// Storage represents all possible actions available to deal with data
type Storage interface {
	SaveBeer(...models.Beer) error
	SaveReview(...models.Review) error
	FindBeer(models.Beer) ([]*models.Beer, error)
	FindReview(models.Review) ([]*models.Review, error)
	FindBeers() []models.Beer
	FindReviews() []models.Review
}

// NewStorage checks storage type (memory or JSON) accordingly creates the DB
func NewStorage(storageType StorageType) error {
	var err error

	switch storageType {
	case MemoryType:
		DB = new(Memory)

	case JSONType:
		// for the moment storage location for JSON files is the current working directory
		DB, err = NewJSON("./data/")
		if err != nil {
			return err
		}
	}

	return nil
}
