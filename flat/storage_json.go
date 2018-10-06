package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/nanobox-io/golang-scribble"
	"github.com/pkg/errors"
)

const (
	// CollectionBeer identifier for JSON collection about beers
	CollectionBeer = "beers"
	// CollectionReview identifier for JSON collection about reviews
	CollectionReview = "reviews"
)

// StorageJSON is the data storage layered using JSON file
type StorageJSON struct {
	db *scribble.Driver
}

func NewStorageJSON(location string) (*StorageJSON, error) {
	var err error

	stg := new(StorageJSON)

	stg.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return stg, nil
}

// SaveBeer saves a new beer if no duplicates have been found.
func (s *StorageJSON) SaveBeer(beers ...Beer) error {
	for _, beer := range beers {
		allBeers, err := s.FindBeers()
		if err != nil {
			return err
		}

		for _, b := range allBeers {
			if beer.Abv == b.Abv &&
				beer.Brewery == b.Brewery &&
				beer.Name == b.Name {
				return fmt.Errorf("beer already exists")
			}
		}

		// for simplicity, and since the delete function does not exist, assume we'll just auto-increment the ID
		beer.ID = len(allBeers) + 1
		beer.Created = time.Now()

		if err := s.db.Write(CollectionBeer, strconv.Itoa(beer.ID), beer); err != nil {
			return err
		}
	}
	return nil
}

// SaveBeer saves a new review.
func (s *StorageJSON) SaveReview(reviews ...Review) error {
	for _, review := range reviews {

		review.Created = time.Now()
		review.ID = fmt.Sprintf("%d_%d", review.BeerID, review.Created.UnixNano())

		if err := s.db.Write(CollectionReview, review.ID, review); err != nil {
			return err
		}
	}

	return nil
}

// FindBeer returns any beers matching the given criteria.
// Beer ID is the only criteria supported at the moment.
func (s *StorageJSON) FindBeer(criteria Beer) ([]Beer, error) {
	var beers []Beer

	if criteria.ID != 0 {
		var beer Beer
		if err := s.db.Read(CollectionBeer, strconv.Itoa(criteria.ID), &beer); err != nil {
			return beers, err
		}

		beers = append(beers, beer)

		return beers, nil
	}

	return beers, fmt.Errorf("no beer ID specified")
}

// FindReview finds all reviews for a given criteria.
// Beer ID is the only criteria supported at the moment.
func (s *StorageJSON) FindReview(criteria Review) ([]Review, error) {
	var matches []Review

	if criteria.BeerID != 0 {
		reviews, err := s.findReviews()
		if err != nil {
			return matches, err
		}

		for _, r := range reviews {
			if r.BeerID == criteria.BeerID {
				matches = append(matches, r)
			}
		}

		return matches, nil
	}

	return matches, fmt.Errorf("no beer ID specified")
}

// findReviews returns all reviews found in the storage
func (s *StorageJSON) findReviews() ([]Review, error) {
	var reviews []Review

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		return reviews, errors.Errorf("failed to fetch all reviews from the JSON storage: %s", err.Error())
	}

	for _, r := range records {
		var review Review

		if err := json.Unmarshal([]byte(r), &review); err != nil {
			return reviews, errors.Errorf("failed to parse review data from the JSON file: %s data: %s", err.Error(), r)
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

// FindBeers returns all beers found in the storage
func (s *StorageJSON) FindBeers() ([]Beer, error) {
	var beers []Beer

	records, err := s.db.ReadAll(CollectionBeer)
	if err != nil {
		return beers, errors.Errorf("failed to fetch all beers from the JSON storage: %s", err.Error())
	}

	for _, b := range records {
		var beer Beer

		if err := json.Unmarshal([]byte(b), &beer); err != nil {
			return beers, errors.Errorf("failed to parse beer data from the JSON file: %s data: %s", err.Error(), b)
		}

		beers = append(beers, beer)
	}

	return beers, nil
}
