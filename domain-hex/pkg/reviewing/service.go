package reviewing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// Repository provides access to the review storage.
type Repository interface {
	// AddReview saves a given review.
	AddReview(Review) error
}

// Service provides reviewing operations.
type Service interface {
	AddBeerReview(Review)
	AddSampleReviews([]Review)
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeerReview saves a new beer review in the database
func (s *service) AddBeerReview(r Review) {
	_ = s.r.AddReview(r) // error handling omitted for simplicity
}

// AddSampleReviews adds some sample reviews to the database
func (s *service) AddSampleReviews(r []Review) {
	for _, rr := range r {
		_ = s.r.AddReview(rr) // error handling omitted for simplicity
	}
}
