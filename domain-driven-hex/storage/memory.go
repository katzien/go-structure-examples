package storage

import (
	"github.com/katzien/structure-examples/domain-driven/beers"
	"github.com/katzien/structure-examples/domain-driven/reviews"
	"time"
	"fmt"
)

// Memory storage keeps beer data in memory
type MemoryBeerStorage struct {
	beers   []beers.Beer
	reviews []reviews.Review
}

// Add saves the given beer to the repository
func (m *MemoryBeerStorage) Add(b beers.Beer) error {
	for _, e := range m.beers {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return beers.ErrDuplicate
		}
	}

	b.ID = len(m.beers) + 1
	b.Created = time.Now()
	m.beers = append(m.beers, b)

	return nil
}

// Get returns a beer with the specified ID
func (m *MemoryBeerStorage) Get(id int) (beers.Beer, error) {
	var beer beers.Beer

	for i := range m.beers {

		if m.beers[i].ID == id {
			return m.beers[i], nil
		}
	}

	return beer, beers.ErrUnknown
}

// GetAll return all beers
func (m *MemoryBeerStorage) GetAll() []beers.Beer {
	return m.beers
}

// Memory storage keeps review data in memory
type MemoryReviewStorage struct {
	beers   []beers.Beer
	reviews []reviews.Review
}

// Add saves the given review in the repository
func (m *MemoryReviewStorage) Add(r reviews.Review) error {
	found := false
	for b := range m.beers {
		if m.beers[b].ID == r.BeerID {
			found = true
		}
	}

	if found {
		r.ID = fmt.Sprintf("%s_%s_%s_%s", r.BeerID, r.FirstName, r.LastName, r.Created.Unix())
		r.Created = time.Now()
		m.reviews = append(m.reviews, r)
	} else {
		return reviews.ErrNotFound
	}

	return nil
}

// GetAll returns all reviews for a given beer
func (m *MemoryReviewStorage) GetAll(beerID int) []reviews.Review {
	var list []reviews.Review

	for i := range m.reviews {
		if m.reviews[i].BeerID == beerID {
			list = append(list, m.reviews[i])
		}
	}

	return list
}
