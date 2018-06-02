package reviewing

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/katzien/structure-examples/domain-driven/reviews"
)

type Handler func(http.ResponseWriter, *http.Request, httprouter.Params)

// MakeAddBeerEndpoint creates a handler for POST /beers/:id/reviews requests
func MakeAddBeerReviewEndpoint(s Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		var newReview reviews.Review
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
