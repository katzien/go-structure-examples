package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/structure-examples/layered/models"
)

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	// TODO: Consider checking if a beer matching the ID actually exists, and
	// 404 if that is not the case.

	results, _ := models.DB.FindReview(models.Review{BeerID: ID})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	var newReview models.Review
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newReview); err != nil {
		http.Error(w, "Failed to parse review", http.StatusBadRequest)
	}

	newReview.BeerID = ID
	if err := models.DB.SaveReview(newReview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New beer review added.")

}
