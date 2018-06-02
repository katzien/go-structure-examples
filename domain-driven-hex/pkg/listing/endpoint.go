package listing

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/katzien/structure-examples/domain-driven/beers"
)

type Handler func(http.ResponseWriter, *http.Request, httprouter.Params)

// MakeAddBeerEndpoint creates a handler for GET /beers requests
func MakeGetBeersEndpoint(s Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetBeers()
		json.NewEncoder(w).Encode(list)
	}
}

// MakeAddBeeEndpoint creates a handler for GET /beers/:id requests
func MakeGetBeerEndpoint(s Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		beer, err := s.GetBeer(ID)
		if err == beers.ErrUnknown {
			http.Error(w, "The beer you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}

// MakeGetBeerReviewsEndpoint creates a handler for GET /beers/:id/reviews requests
func MakeGetBeerReviewsEndpoint(s Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		reviews, err := s.GetBeerReviews(ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
