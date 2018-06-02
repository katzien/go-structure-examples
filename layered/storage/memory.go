package storage

import "github.com/katzien/structure-examples/layered/models"

// Memory data storage layered save only in memory
type Memory struct {
	cellar  []models.Beer
	reviews []models.Review
}

// SaveBeer insert or update beers
func (s *Memory) SaveBeer(beers ...models.Beer) error {
	for _, beer := range beers {
		var err error

		beersFound, err := s.FindBeer(beer)
		if err != nil {
			return err
		}

		if len(beersFound) == 1 {
			*beersFound[0] = beer
			return nil
		}

		beer.ID = len(s.cellar) + 1
		s.cellar = append(s.cellar, beer)
	}

	return nil
}

// SaveReview insert or update reviews
func (s *Memory) SaveReview(reviews ...models.Review) error {
	for _, review := range reviews {
		var err error

		reviewsFound, err := s.FindReview(review)
		if err != nil {
			return err
		}

		if len(reviewsFound) == 1 {
			*reviewsFound[0] = review
			return nil
		}

		review.ID = len(s.reviews) + 1
		s.reviews = append(s.reviews, review)
	}

	return nil
}

// FindBeer locate full data set based on given criteria
func (s *Memory) FindBeer(criteria models.Beer) ([]*models.Beer, error) {
	var beers []*models.Beer

	for idx := range s.cellar {

		if s.cellar[idx].ID == criteria.ID {
			beers = append(beers, &s.cellar[idx])
		}
	}

	return beers, nil
}

// FindReview locate full data set based on given criteria
func (s *Memory) FindReview(criteria models.Review) ([]*models.Review, error) {
	var reviews []*models.Review

	for idx := range s.reviews {
		if s.reviews[idx].ID == criteria.ID || s.reviews[idx].BeerID == criteria.BeerID {
			reviews = append(reviews, &s.reviews[idx])
		}
	}

	return reviews, nil
}

// FindBeers return all beers
func (s *Memory) FindBeers() []models.Beer {
	return s.cellar
}

// FindReviews return all reviews
func (s *Memory) FindReviews() []models.Review {
	return s.reviews
}
