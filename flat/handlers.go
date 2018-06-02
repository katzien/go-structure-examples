package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetBeers returns the cellar
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	cellar := db.FindBeers()
	json.NewEncoder(w).Encode(cellar)
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	cellar, _ := db.FindBeer(Beer{ID: ID})
	if len(cellar) == 1 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cellar[0])
		return
	}

	http.Error(w, "The beer you requested does not exist.", http.StatusNotFound)
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	// TODO: Consider checking if a beer matching the ID actually exists, and
	// 404 if that is not the case.

	results, _ := db.FindReview(Review{BeerID: ID})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// AddBeer adds a new beer to the cellar
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var newBeer Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		http.Error(w, "Bad beer - this will be a HTTP status code soon!", http.StatusBadRequest)
		return
	}

	db.SaveBeer(newBeer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New beer added.")
}

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	var newReview Review
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newReview); err != nil {
		http.Error(w, "Failed to parse review", http.StatusBadRequest)
	}

	newReview.BeerID = ID
	if err := db.SaveReview(newReview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New beer review added.")

}
