package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/listing"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/reviewing"
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New beer added.")
	}
}

// addBeerReview returns a handler for POST /beers/:id/reviews requests
func addBeerReview(s reviewing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		var newReview reviewing.Review
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&newReview); err != nil {
			http.Error(w, "Failed to parse review", http.StatusBadRequest)
		}

		newReview.BeerID = ID

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
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		beer, err := s.GetBeer(ID)
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
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		reviews := s.GetBeerReviews(ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
