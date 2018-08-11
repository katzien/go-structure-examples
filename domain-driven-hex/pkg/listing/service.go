package listing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// Repository provides access to the beer and review data.
type Repository interface {
	GetBeer(beerID int) (Beer, error)
	GetAllBeers() []Beer
	GetAllReviews(beerID int) ([]Review)
}

// Service provides beer or review adding operations
type Service interface {
	GetBeer(int) (Beer, error)
	GetBeers() []Beer
	GetBeerReviews(int) ([]Review)
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
func (s *service) GetBeer(id int) (Beer, error) {
	return s.r.GetBeer(id)
}

// GetBeerReviews returns all requests for a beer
func (s *service) GetBeerReviews(beerID int) ([]Review) {
	return s.r.GetAllReviews(beerID)
}