package json

import (
	"encoding/json"
	"fmt"

	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/listing"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/reviewing"
	"github.com/nanobox-io/golang-scribble"
	"path"
	"runtime"
	"strconv"
	"time"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionBeer identifier for the JSON collection of beers
	CollectionBeer = "beers"
	// CollectionReview identifier for the JSON collection of reviews
	CollectionReview = "reviews"
)

// Storage stores beer data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddBeer saves the given beer to the repository
func (s *Storage) AddBeer(b adding.Beer) error {

	existingBeers := s.GetAllBeers()
	for _, e := range existingBeers {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newB := Beer{
		ID:        len(existingBeers) + 1,
		Created:   time.Now(),
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
	}

	resource := strconv.Itoa(newB.ID)
	if err := s.db.Write(CollectionBeer, resource, newB); err != nil {
		return err
	}
	return nil
}

// AddReview saves the given review in the repository
func (s *Storage) AddReview(r reviewing.Review) error {

	var beer Beer
	if err := s.db.Read(CollectionBeer, strconv.Itoa(r.BeerID), &beer); err != nil {
		return listing.ErrNotFound
	}

	created := time.Now()
	newR := Review{
		ID:        fmt.Sprintf("%d_%s_%s_%d", r.BeerID, r.FirstName, r.LastName, created.Unix()),
		Created:   created,
		BeerID:    r.BeerID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
	}

	if err := s.db.Write(CollectionReview, newR.ID, r); err != nil {
		return err
	}

	return nil
}

// Get returns a beer with the specified ID
func (s *Storage) GetBeer(id int) (listing.Beer, error) {
	var b Beer
	var beer listing.Beer

	var resource = strconv.Itoa(id)

	if err := s.db.Read(CollectionBeer, resource, &b); err != nil {
		// err handling omitted for simplicity
		return beer, listing.ErrNotFound
	}

	beer.ID = b.ID
	beer.Name = b.Name
	beer.Brewery = b.Brewery
	beer.Abv = b.Abv
	beer.ShortDesc = b.ShortDesc
	beer.Created = b.Created

	return beer, nil
}

// GetAll returns all beers
func (s *Storage) GetAllBeers() []listing.Beer {
	list := []listing.Beer{}

	records, err := s.db.ReadAll(CollectionBeer)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var b Beer
		var beer listing.Beer

		if err := json.Unmarshal([]byte(r), &b); err != nil {
			// err handling omitted for simplicity
			return list
		}

		beer.ID = b.ID
		beer.Name = b.Name
		beer.Brewery = b.Brewery
		beer.Abv = b.Abv
		beer.ShortDesc = b.ShortDesc
		beer.Created = b.Created

		list = append(list, beer)
	}

	return list
}

// GetAll returns all reviews for a given beer
func (s *Storage) GetAllReviews(beerID int) []listing.Review {
	list := []listing.Review{}

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, b := range records {
		var r Review

		if err := json.Unmarshal([]byte(b), &r); err != nil {
			// err handling omitted for simplicity
			return list
		}

		if r.BeerID == beerID {
			var review listing.Review

			review.ID = r.ID
			review.BeerID = r.BeerID
			review.FirstName = r.FirstName
			review.LastName = r.LastName
			review.Score = r.Score
			review.Text = r.Text
			review.Created = r.Created

			list = append(list, review)
		}
	}

	return list
}
