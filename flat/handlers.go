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
	cellar, err := db.FindBeers()
	if err != nil {
		// TODO: return the right HTTP status based on the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(cellar)
	return
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	result, err := db.FindBeer(Beer{ID: ID})
	if err != nil {
		// TODO: return the right HTTP status based on the error
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	result, err := db.FindReview(Review{BeerID: ID})
	if err != nil {
		// TODO: return the right HTTP status based on the error
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return
}

// AddBeer adds a new beer to the cellar
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var newBeer Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		http.Error(w, "error while parsing new beer data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.SaveBeer(newBeer); err != nil {
		// TODO: return the right HTTP status based on the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	var newReview Review
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newReview); err != nil {
		http.Error(w, "error while parsing new review data: "+err.Error(), http.StatusBadRequest)
		return
	}

	newReview.BeerID = ID
	if err := db.SaveReview(newReview); err != nil {
		// TODO: return the right HTTP status based on the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}
