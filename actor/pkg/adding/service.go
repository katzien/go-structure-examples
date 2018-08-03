package adding

import (
	"github.com/katzien/go-structure-examples/actor/pkg/beers"
)

type Payload []beers.Beer

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

// Service provides beer or review adding operations
type Service interface {
	AddBeer(b ...beers.Beer)
	AddSampleBeers(chan<- Event, Payload)
}

type service struct {
	bR beers.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(bR beers.Repository) Service {
	return &service{bR}
}

// AddBeer adds the given beer(s) to the database
func (s *service) AddBeer(b ...beers.Beer) {
	for _, beer := range b {
		_ = s.bR.Add(beer) // error handling omitted for simplicity
	}
}

// AddSampleBeers adds some sample beers to the database
func (s *service) AddSampleBeers(done chan<- Event, data Payload) {
	for _, b := range data {
		err := s.bR.Add(b)
		if err != nil {
			if err == beers.ErrDuplicate {
				// forgive the naughty error type checking above...
				done <- BeerAlreadyExists
			}
			done <- Failed
		}
	}

	done <- Done
}