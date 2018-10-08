package main

import (
	"fmt"
	"time"
)

// StorageMemory data storage layered save only in memory
type StorageMemory struct {
	cellar  []Beer
	reviews []Review
}

// SaveBeer insert or update beers
func (s *StorageMemory) SaveBeer(beers ...Beer) error {
	for _, beer := range beers {

		for _, b := range s.cellar {
			if beer.Abv == b.Abv &&
				beer.Brewery == b.Brewery &&
				beer.Name == b.Name {
				return fmt.Errorf("beer already exists")
			}
		}

		// for simplicity, and since the delete function does not exist, assume we'll just auto-increment the ID
		beer.ID = len(s.cellar) + 1
		beer.Created = time.Now()

		s.cellar = append(s.cellar, beer)
	}

	return nil
}

// SaveReview insert or update reviews
func (s *StorageMemory) SaveReview(reviews ...Review) error {
	for _, review := range reviews {

		review.ID = fmt.Sprintf("%d_%d", review.BeerID, review.Created.Unix())
		review.Created = time.Now()

		s.reviews = append(s.reviews, review)
	}

	return nil
}

// FindBeer returns any beers matching the given criteria.
// Beer ID is the only criteria supported at the moment.
func (s *StorageMemory) FindBeer(criteria Beer) ([]Beer, error) {
	var beers []Beer

	for _, b := range s.cellar {

		if b.ID == criteria.ID {
			beers = append(beers, b)
		}
	}

	return beers, nil
}

// FindReview finds all reviews for a given criteria.
// Beer ID is the only criteria supported at the moment.
func (s *StorageMemory) FindReview(criteria Review) ([]Review, error) {
	var matches []Review

	for _, r := range s.reviews {
		if r.BeerID == criteria.BeerID {
			matches = append(matches, r)
		}
	}

	return matches, nil
}

// FindBeers return all beers
func (s *StorageMemory) FindBeers() ([]Beer, error) {
	return s.cellar, nil
}
