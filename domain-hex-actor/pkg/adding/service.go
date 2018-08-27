package adding

import (
	"errors"
)

type Payload []Beer

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota

	// BeerAlreadyExists means the given beer is a duplicate of an existing one
	BeerAlreadyExists

	// Failed means processing did not finish successfully
	Failed
)

func (e Event) GetMeaning() string {
	if e == Done {
		return "Done"
	}

	if e == BeerAlreadyExists {
		return "Duplicate beer"
	}

	if e == Failed {
		return "Failed"
	}

	return "Unknown result"
}

var ErrDuplicate = errors.New("beer already exists")

// Service provides beer adding operations.
type Service interface {
	AddBeer(...Beer)
	AddSampleBeers(Payload) <-chan Event
}

// Repository provides access to beer repository.
type Repository interface {
	// AddBeer saves a given beer to the repository.
	AddBeer(Beer) error
}

type service struct {
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeer adds the given beer(s) to the database
func (s *service) AddBeer(b ...Beer) {

	// any validation can be done here

	for _, beer := range b {
		_ = s.bR.AddBeer(beer) // error handling omitted for simplicity
	}
}

// AddSampleBeers adds some sample beers to the database
func (s *service) AddSampleBeers(data Payload) <-chan Event {
	results := make(chan Event)

	go func() {
		defer close(results)

		for _, b := range data {
			err := s.bR.AddBeer(b)
			if err != nil {
				if err == ErrDuplicate {
					// forgive the naughty error type checking above...
					results <- BeerAlreadyExists
					continue
				}
				results <- Failed
				continue
			}

			results <- Done
		}
	}()

	return results
}
