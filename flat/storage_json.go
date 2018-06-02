package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nanobox-io/golang-scribble"
)

const (
	// CollectionBeer identifier for JSON collection about beers
	CollectionBeer int = iota
	// CollectionReview identifier for JSON collection about reviews
	CollectionReview
)

// StorageJSON is the data storage layered using JSON file
type StorageJSON struct {
	db *scribble.Driver
}

func newStorageJSON(location string) (*StorageJSON, error) {
	var err error

	stg := new(StorageJSON)

	stg.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return stg, nil
}

// SaveBeer insert new beers
func (s *StorageJSON) SaveBeer(beers ...Beer) error {
	for _, beer := range beers {
		var resource = strconv.Itoa(beer.ID)
		var collection = strconv.Itoa(CollectionBeer)

		allBeers := s.FindBeers()
		for _, b := range allBeers {
			if beer.Abv == b.Abv &&
				beer.Brewery == b.Brewery &&
				beer.Name == b.Name {
				return fmt.Errorf("Beer already exists")
			}
		}

		// TODO: Since delete function has not been implemented yet
		// I think we can assume size of beers should always increase.
		beer.ID = len(allBeers) + 1

		if err := s.db.Write(collection, resource, beer); err != nil {
			return err
		}
	}
	return nil
}

// SaveReview insert reviews
func (s *StorageJSON) SaveReview(reviews ...Review) error {
	for _, review := range reviews {
		var resource = strconv.Itoa(review.ID)
		var collection = strconv.Itoa(CollectionReview)

		beerFound, err := s.FindBeer(Beer{ID: review.BeerID})
		if err != nil {
			return err
		}

		if len(beerFound) == 0 {
			return fmt.Errorf("The beer selected for the review does not exist")
		}

		allReviews := s.FindReviews()
		for _, r := range allReviews {
			if review.BeerID == r.BeerID &&
				review.FirstName == r.FirstName &&
				review.LastName == r.LastName &&
				review.Text == r.Text {
				return fmt.Errorf("Review already exists")
			}
		}

		// TODO: Since delete function has not been implemented yet
		// I think we can assume size of reviews should always increase.
		review.ID = len(allReviews) + 1

		if err = s.db.Write(collection, resource, review); err != nil {
			return err
		}
	}
	return nil
}

// FindBeer locate full data set based on given criteria
func (s *StorageJSON) FindBeer(criteria Beer) ([]*Beer, error) {
	var beers []*Beer
	var beer Beer
	var resource = strconv.Itoa(criteria.ID)
	var collection = strconv.Itoa(CollectionBeer)

	if err := s.db.Read(collection, resource, &beer); err != nil {
		return beers, err
	}

	beers = append(beers, &beer)

	return beers, nil
}

// FindReview locate full data set based on given criteria
func (s *StorageJSON) FindReview(criteria Review) ([]*Review, error) {
	var reviews []*Review
	var review Review
	var resource = strconv.Itoa(criteria.ID)
	var collection = strconv.Itoa(CollectionReview)

	if err := s.db.Read(collection, resource, &review); err != nil {
		return reviews, err
	}

	reviews = append(reviews, &review)

	return reviews, nil
}

func (s *StorageJSON) FindBeers() []Beer {
	var beers []Beer
	var collection = strconv.Itoa(CollectionBeer)

	records, err := s.db.ReadAll(collection)
	if err != nil {
		return beers
	}

	for _, b := range records {
		var beer Beer

		if err := json.Unmarshal([]byte(b), &beer); err != nil {
			return beers
		}

		beers = append(beers, beer)
	}

	return beers
}

func (s *StorageJSON) FindReviews() []Review {
	var reviews []Review
	var collection = strconv.Itoa(CollectionReview)

	records, err := s.db.ReadAll(collection)
	if err != nil {
		return reviews
	}

	for _, r := range records {
		var review Review

		if err := json.Unmarshal([]byte(r), &review); err != nil {
			return reviews
		}

		reviews = append(reviews, review)
	}

	return reviews
}
