package listing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetBeer returns the beer with given ID.
	GetBeer(string) (Beer, error)
	// GetAllBeers returns all beers saved in storage.
	GetAllBeers() []Beer
	// GetAllReviews returns a list of all reviews for a given beer ID.
	GetAllReviews(string) []Review
}

// Service provides beer and review listing operations.
type Service interface {
	GetBeer(string) (Beer, error)
	GetBeers() []Beer
	GetBeerReviews(string) []Review
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) GetBeers() []Beer {
	return s.r.GetAllBeers()
}

// GetBeer returns a beer
func (s *service) GetBeer(id string) (Beer, error) {
	return s.r.GetBeer(id)
}

// GetBeerReviews returns all requests for a beer
func (s *service) GetBeerReviews(beerID string) []Review {
	return s.r.GetAllReviews(beerID)
}
