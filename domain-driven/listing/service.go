package listing

import (
	"github.com/katzien/structure-examples/domain-driven/reviews"
	"github.com/katzien/structure-examples/domain-driven/beers"
)

// Service provides beer or review adding operations
type Service interface {
	GetBeers() []beers.Beer
	GetBeer(int) (beers.Beer, error)
	GetBeerReviews(int) ([]reviews.Review, error)
}

type service struct {
	bR beers.Repository
	rR reviews.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(bR beers.Repository, rR reviews.Repository) Service {
	return &service{bR, rR}
}

// GetBeers returns all beers
func (s *service) GetBeers() []beers.Beer {
	return s.bR.GetAll()
}

// GetBeer returns a beer
func (s *service) GetBeer(id int) (beers.Beer, error) {
	return s.bR.Get(id)
}

// GetBeerReviews returns all requests for a beer
func (s *service) GetBeerReviews(beerID int) ([]reviews.Review, error) {
	var list []reviews.Review
	if _, err := s.bR.Get(beerID); err == beers.ErrUnknown {
		return list, reviews.ErrNotFound
	}

	return s.rR.GetAll(beerID), nil
}