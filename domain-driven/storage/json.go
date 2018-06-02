package storage

import (
	"encoding/json"
	"fmt"

	"github.com/nanobox-io/golang-scribble"
	"strconv"
	"github.com/katzien/structure-examples/domain-driven/beers"
	"github.com/katzien/structure-examples/domain-driven/reviews"
	"time"
)

const (
	// location defines where the files are stored
	location = "../../storage/json/"

	// CollectionBeer identifier for the JSON collection of beers
	CollectionBeer = "beers"
	// CollectionReview identifier for the JSON collection of reviews
	CollectionReview = "reviews"
)

// JSONBeerStorage stores beer data in JSON files
type JSONBeerStorage struct {
	db *scribble.Driver
}

// NewJSONBeerStorage returns a new JSON beer storage
func NewJSONBeerStorage() (*JSONBeerStorage, error) {
	var err error

	s := new(JSONBeerStorage)

	s.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Add saves the given beer to the repository
func (s *JSONBeerStorage) Add(b beers.Beer) error {
	var resource = strconv.Itoa(b.ID)

	existingBeers := s.GetAll()
	for _, e := range existingBeers {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return beers.ErrDuplicate
		}
	}

	b.ID = len(existingBeers) + 1
	b.Created = time.Now()

	if err := s.db.Write(CollectionBeer, resource, b); err != nil {
		return err
	}
	return nil
}

// Get returns a beer with the specified ID
func (s *JSONBeerStorage) Get(id int) (beers.Beer, error) {
	var beer beers.Beer
	var resource = strconv.Itoa(id)

	if err := s.db.Read(CollectionBeer, resource, &beer); err != nil {
		return beer, beers.ErrUnknown
	}

	return beer, nil
}

// GetAll returns all beers
func (s *JSONBeerStorage) GetAll() []beers.Beer {
	var list []beers.Beer

	records, err := s.db.ReadAll(CollectionBeer)
	if err != nil {
		panic("error while fetching beers from the JSON file storage: %v" + err.Error())
	}

	for _, b := range records {
		var beer beers.Beer

		if err := json.Unmarshal([]byte(b), &beer); err != nil {
			return list
		}

		list = append(list, beer)
	}

	return list
}

// JSONReviewStorage stores review data in JSON files
type JSONReviewStorage struct {
	db *scribble.Driver
}

// NewJSONReviewStorage returns a new JSON reviews storage
func NewJSONReviewStorage() (*JSONReviewStorage, error) {
	var err error

	s := new(JSONReviewStorage)

	s.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Add saves the given review in the repository
func (s *JSONReviewStorage) Add(r reviews.Review) error {

	var beer beers.Beer
	if err := s.db.Read(CollectionBeer, strconv.Itoa(r.BeerID), &beer); err != nil {
		return reviews.ErrNotFound
	}

	r.ID = fmt.Sprintf("%s_%s_%s_%s", r.BeerID, r.FirstName, r.LastName, r.Created.Unix())
	r.Created = time.Now()

	if err := s.db.Write(CollectionReview, r.ID, r); err != nil {
		return err
	}

	return nil
}

// GetAll returns all reviews for a given beer
func (s *JSONReviewStorage) GetAll(beerID int) []reviews.Review {
	var list []reviews.Review

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		panic("error while fetching reviews from the JSON file storage: " + err.Error())
	}

	for _, r := range records {
		var review reviews.Review

		if err := json.Unmarshal([]byte(r), &review); err != nil {
			return list
		}

		list = append(list, review)
	}

	return list
}
