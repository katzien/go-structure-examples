package storage

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/katzien/go-structure-examples/modular/beers"
	"github.com/katzien/go-structure-examples/modular/reviews"
	"github.com/nanobox-io/golang-scribble"
)

const (
	// CollectionBeer identifier for JSON collection about beers
	CollectionBeer int = iota
	// CollectionReview identifier for JSON collection about reviews
	CollectionReview
)

// JSON is the data storage layered using JSON file
type JSONStorage struct {
	db *scribble.Driver
}

func NewJSONStorage(location string) (*JSONStorage, error) {
	var err error

	stg := new(JSONStorage)

	stg.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return stg, nil
}

// SaveBeer insert new beers
func (s *JSONStorage) SaveBeer(beers ...beers.Beer) error {
	for _, beer := range beers {
		var resource = strconv.Itoa(beer.ID)
		var collection = strconv.Itoa(CollectionBeer)

		allBeers := s.FindBeers()
		for _, b := range allBeers {
			if beer.Abv == b.Abv &&
				beer.Brewery == b.Brewery &&
				beer.Name == b.Name {
				return fmt.Errorf("beer already exists")
			}
		}

		// for simplicity we'll assume the IDs will always increase,
		// since there's no delete functionality
		beer.ID = len(allBeers) + 1

		if err := s.db.Write(collection, resource, beer); err != nil {
			return err
		}
	}
	return nil
}

// SaveReview insert reviews
func (s *JSONStorage) SaveReview(reviews ...reviews.Review) error {
	for _, review := range reviews {
		var collection = strconv.Itoa(CollectionReview)

		beerFound, err := s.FindBeer(beers.Beer{ID: review.BeerID})
		if err != nil {
			return err
		}

		if len(beerFound) == 0 {
			return fmt.Errorf("beer selected for review does not exist")
		}

		allReviews := s.FindReviews()
		for _, r := range allReviews {
			if review.BeerID == r.BeerID &&
				review.FirstName == r.FirstName &&
				review.LastName == r.LastName &&
				review.Text == r.Text {
				return fmt.Errorf("review already exists")
			}
		}

		review.ID = fmt.Sprintf("%d_%s_%s_%d", review.BeerID, review.FirstName, review.LastName, time.Now().Unix())

		if err = s.db.Write(collection, review.ID, review); err != nil {
			return err
		}
	}
	return nil
}

// FindBeer locate full data set based on given criteria
func (s *JSONStorage) FindBeer(criteria beers.Beer) ([]*beers.Beer, error) {
	var beers []*beers.Beer
	var beer beers.Beer
	var resource = strconv.Itoa(criteria.ID)
	var collection = strconv.Itoa(CollectionBeer)

	if err := s.db.Read(collection, resource, &beer); err != nil {
		return beers, err
	}

	beers = append(beers, &beer)

	return beers, nil
}

// FindReview locate full data set based on given criteria
func (s *JSONStorage) FindReview(criteria reviews.Review) ([]*reviews.Review, error) {
	var reviews []*reviews.Review
	var review reviews.Review
	var resource = strconv.Itoa(criteria.ID)
	var collection = strconv.Itoa(CollectionReview)

	if err := s.db.Read(collection, resource, &review); err != nil {
		return reviews, err
	}

	reviews = append(reviews, &review)

	return reviews, nil
}

func (s *JSONStorage) FindBeers() []beers.Beer {
	var beers []beers.Beer
	var collection = strconv.Itoa(CollectionBeer)

	records, err := s.db.ReadAll(collection)
	if err != nil {
		return beers
	}

	for _, b := range records {
		var beer beers.Beer

		if err := json.Unmarshal([]byte(b), &beer); err != nil {
			return beers
		}

		beers = append(beers, beer)
	}

	return beers
}

func (s *JSONStorage) FindReviews() []reviews.Review {
	var reviews []reviews.Review
	var collection = strconv.Itoa(CollectionReview)

	records, err := s.db.ReadAll(collection)
	if err != nil {
		return reviews
	}

	for _, r := range records {
		var review reviews.Review

		if err := json.Unmarshal([]byte(r), &review); err != nil {
			return reviews
		}

		reviews = append(reviews, review)
	}

	return reviews
}
