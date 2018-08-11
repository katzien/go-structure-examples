package reviewing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// Repository provides access to the reviews.
type Repository interface {
	AddReview(review Review) error
}

// Service provides beer or review adding operations
type Service interface {
	AddBeerReview(r Review)
	AddSampleReviews()
}

type service struct {
	rR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeerReview saves a new beer review in the database
func (s *service) AddBeerReview(r Review) {
	_ = s.rR.AddReview(r) // error handling omitted for simplicity
}

// AddSampleReviews adds some sample reviews to the database
func (s *service) AddSampleReviews() {
	for _, b := range DefaultReviews {
		_ = s.rR.AddReview(b) // error handling omitted for simplicity
	}
}