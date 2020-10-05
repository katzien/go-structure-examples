package adding

import (
	"errors"
	"github.com/katzien/go-structure-examples/domain-hex/pkg/listing"
)

// ErrDuplicate is used when a beer already exists.
var ErrDuplicate = errors.New("beer already exists")

// Service provides beer adding operations.
type Service interface {
	AddBeer(...Beer) error
	AddSampleBeers([]Beer)
}

// Repository provides access to beer repository.
type Repository interface {
	// AddBeer saves a given beer to the repository.
	AddBeer(Beer) error
	// GetAllBeers returns all beers saved in storage.
	GetAllBeers() []listing.Beer
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeer persists the given beer(s) to storage
func (s *service) AddBeer(b ...Beer) error {
	// make sure we don't add any duplicates
	existingBeers := s.r.GetAllBeers()
	for _, bb := range b {
		for _, e := range existingBeers {
			if bb.Abv == e.Abv &&
				bb.Brewery == e.Brewery &&
				bb.Name == e.Name {
				return ErrDuplicate
			}
		}
	}

	// any other validation can be done here

	for _, beer := range b {
		_ = s.r.AddBeer(beer) // error handling omitted for simplicity
	}

	return nil
}

// AddSampleBeers adds some sample beers to the database
func (s *service) AddSampleBeers(b []Beer) {

	// any validation can be done here

	for _, bb := range b {
		_ = s.r.AddBeer(bb) // error handling omitted for simplicity
	}
}
