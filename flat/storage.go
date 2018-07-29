package main

// StorageType defines available storage types
type StorageType int

const (
	// JSONDataLocation defines where the JSON data files will be saved
	JSONDataLocation = "./data/"

	// JSON will store data in JSON files saved on disk
	JSON StorageType = iota
	// Memory will store data in memory
	Memory
)

// Storage represents all possible actions available to deal with data
type Storage interface {
	SaveBeer(...Beer) error
	SaveReview(...Review) error
	FindBeer(Beer) ([]Beer, error)
	FindReview(Review) ([]Review, error)
	FindBeers() ([]Beer, error)
}

func NewStorage(storageType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case Memory:
		stg = new(StorageMemory)

	case JSON:
		// for the moment storage location for JSON files is the current working directory
		stg, err = NewStorageJSON(JSONDataLocation)
	}

	return stg, err
}
