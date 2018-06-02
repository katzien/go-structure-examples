package reviewing

import (
	"github.com/katzien/structure-examples/domain-driven/reviews"
)

// Service provides beer or review adding operations
type Service interface {
	AddBeerReview(r reviews.Review)
	AddSampleReviews()
}

type service struct {
	rR reviews.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(rR reviews.Repository) Service {
	return &service{rR}
}

// AddBeerReview saves a new beer review in the database
func (s *service) AddBeerReview(r reviews.Review) {
	_ = s.rR.Add(r) // error handling omitted for simplicity
}

// AddSampleReviews adds some sample reviews to the database
func (s *service) AddSampleReviews() {
	for _, b := range reviews.DefaultReviews {
		_ = s.rR.Add(b) // error handling omitted for simplicity
	}
}