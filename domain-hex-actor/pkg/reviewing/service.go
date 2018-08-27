package reviewing

type Payload []Review

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota

	// Failed means processing did not finish successfully
	Failed

	// We could also have a Queued Event which would mean queued for processing
	Queued
)

// Repository provides access to the review storage.
type Repository interface {
	// AddReview saves a given review.
	AddReview(Review) error
}

// Service provides reviewing operations.
type Service interface {
	AddBeerReview(Review)
	AddSampleReviews(Payload) <-chan Event
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
func (s *service) AddSampleReviews(data Payload) <-chan Event {
	results := make(chan Event)

	go func() {
		defer close(results)

		for _, b := range data {
			err := s.rR.AddReview(b)
			if err != nil {
				results <- Failed
			}
			results <- Done
		}
	}()

	return results
}
