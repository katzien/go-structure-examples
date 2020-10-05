package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/go-structure-examples/domain-hex/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex/pkg/listing"
	"github.com/katzien/go-structure-examples/domain-hex/pkg/reviewing"
)

func Handler(a adding.Service, l listing.Service, r reviewing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", getBeers(l))
	router.GET("/beers/:id", getBeer(l))
	router.GET("/beers/:id/reviews", getBeerReviews(l))

	router.POST("/beers", addBeer(a))
	router.POST("/beers/:id/reviews", addBeerReview(r))

	return router
}

// addBeer returns a handler for POST /beers requests
func addBeer(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newBeer adding.Beer
		err := decoder.Decode(&newBeer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddBeer(newBeer)
		// error handling omitted for simplicity

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New beer added.")
	}
}

// addBeerReview returns a handler for POST /beers/:id/reviews requests
func addBeerReview(s reviewing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var newReview reviewing.Review
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&newReview); err != nil {
			http.Error(w, "Failed to parse review", http.StatusBadRequest)
		}

		newReview.BeerID = p.ByName("id")

		s.AddBeerReview(newReview)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New beer review added.")
	}
}

// getBeers returns a handler for GET /beers requests
func getBeers(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetBeers()
		json.NewEncoder(w).Encode(list)
	}
}

// getBeer returns a handler for GET /beers/:id requests
func getBeer(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		beer, err := s.GetBeer(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The beer you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}

// getBeerReviews returns a handler for GET /beers/:id/reviews requests
func getBeerReviews(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		reviews := s.GetBeerReviews(p.ByName("id"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
