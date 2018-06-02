package handlers

import (
	"net/http/httptest"
	"github.com/vulcand/vulcand/router"
	"bytes"
	"fmt"
	"testing"
	"github.com/katzien/structure-examples/layered/models"
	"encoding/json"
	"math/rand"
	"net/http"
)

//func TestGetBeers(t *testing.T) {
//	var cellarFromRequest []models.Beer
//	var cellarFromStorage []models.Beer
//
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("GET", "/beers", nil)
//
//	router.ServeHTTP(w, r)
//
//	cellarFromStorage = models.DB.FindBeers()
//	json.Unmarshal(w.Body.Bytes(), &cellarFromRequest)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected route GET /beers to be valid.")
//		t.FailNow()
//	}
//
//	if len(cellarFromRequest) != len(cellarFromStorage) {
//		t.Error("Expected number of beers from request to be the same as beers in the storage")
//		t.FailNow()
//	}
//
//	var mapCellar = make(map[models.Beer]int, len(cellarFromStorage))
//	for _, beer := range cellarFromStorage {
//		mapCellar[beer] = 1
//	}
//
//	for _, beerResp := range cellarFromRequest {
//		if _, ok := mapCellar[beerResp]; !ok {
//			t.Errorf("Expected all results to match existing records")
//			t.FailNow()
//			break
//		}
//	}
//}
//
//func TestAddBeer(t *testing.T) {
//	newBeer := models.Beer{
//		Name:    "Testing beer",
//		Abv:     333,
//		Brewery: "Testing Beer Inc",
//	}
//
//	newBeerJSON, err := json.Marshal(newBeer)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("POST", "/beers", bytes.NewBuffer(newBeerJSON))
//
//	router.ServeHTTP(w, r)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected route POST /beers to be valid.")
//		t.FailNow()
//	}
//
//	newBeerMissing := true
//	for _, b := range models.DB.FindBeers() {
//		if b.Name == newBeer.Name &&
//			b.Abv == newBeer.Abv &&
//			b.Brewery == newBeer.Brewery {
//			newBeerMissing = false
//		}
//	}
//
//	if newBeerMissing {
//		t.Errorf("Expected to find new entry in storage`")
//		t.FailNow()
//	}
//}
//
//func TestGetBeer(t *testing.T) {
//	cellar := models.DB.FindBeers()
//	choice := rand.Intn(len(cellar) - 1)
//
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("GET", fmt.Sprintf("/beers/%d", cellar[choice].ID), nil)
//
//	router.ServeHTTP(w, r)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected route GET /beers/%d to be valid.", cellar[choice].ID)
//		t.FailNow()
//	}
//
//	var selectedBeer models.Beer
//	json.Unmarshal(w.Body.Bytes(), &selectedBeer)
//
//	if cellar[choice] != selectedBeer {
//		t.Errorf("Expected to match results with selected beer")
//		t.FailNow()
//	}
//}
