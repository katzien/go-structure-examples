package reviewing

import (
	"github.com/katzien/go-structure-examples/actor/pkg/reviews"
)

type Payload []reviews.Review

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota

	// Queued means queued for processing
	Queued

	// Failed means processing did not finish successfully
	Failed
)

// Service provides beer or review adding operations
type Service interface {
	AddBeerReview(r reviews.Review)
	AddSampleReviews(chan <-Event, Payload)
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
func (s *service) AddSampleReviews(done chan<- Event, data Payload) {
	for _, b := range data {
		err := s.rR.Add(b)
		if err != nil {
			done <- Failed
		}
	}

	done <- Done
}