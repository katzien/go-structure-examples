package adding

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"github.com/katzien/structure-examples/domain-driven/beers"
)

// MakeAddBeerEndpoint creates a handler for POST /beers requests
func MakeAddBeerEndpoint(s Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newBeer beers.Beer
		err := decoder.Decode(&newBeer)
		if err != nil {
			http.Error(w, "Bad beer - this will be a HTTP status code soon!", http.StatusBadRequest)
			return
		}

		s.AddBeer(newBeer)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New beer added.")
	}
}
