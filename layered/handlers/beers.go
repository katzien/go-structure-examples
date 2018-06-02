package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/structure-examples/layered/models"
)

// GetBeers returns the cellar
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	cellar := models.DB.FindBeers()
	json.NewEncoder(w).Encode(cellar)
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	cellar, _ := models.DB.FindBeer(models.Beer{ID: ID})
	if len(cellar) == 1 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cellar[0])
		return
	}

	http.Error(w, "The beer you requested does not exist.", http.StatusNotFound)
}

// AddBeer adds a new beer to the cellar
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var newBeer models.Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		http.Error(w, "Bad beer - this will be a HTTP status code soon!", http.StatusBadRequest)
		return
	}

	models.DB.SaveBeer(newBeer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New beer added.")
}
