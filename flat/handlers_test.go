package main

//func TestGetBeers(t *testing.T) {
//
//	var cellarFromRequest []Beer
//	var cellarFromStorage = db.FindBeers()
//
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("GET", "/beers", nil)
//
//	router.ServeHTTP(w, r)
//
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
//	var mapCellar = make(map[Beer]int, len(cellarFromStorage))
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
//func TestGetBeer(t *testing.T) {
//	cellar := db.FindBeers()
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
//	var selectedBeer Beer
//	json.Unmarshal(w.Body.Bytes(), &selectedBeer)
//
//	if cellar[choice] != selectedBeer {
//		t.Errorf("Expected to match results with selected beer")
//		t.FailNow()
//	}
//
//}
//
//func TestAddBeer(t *testing.T) {
//	newBeer := Beer{
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
//	for _, b := range db.FindBeers() {
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
//
//}
//
//func TestAddBeerReview(t *testing.T) {
//	newReview := Review{
//		BeerID:    1,
//		FirstName: "John",
//		LastName:  "Gopher",
//		Score:     8,
//		Text:      "Decent beer.",
//	}
//
//	newReviewJSON, err := json.Marshal(newReview)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("POST", "/beers/1/reviews", bytes.NewBuffer(newReviewJSON))
//
//	router.ServeHTTP(w, r)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected route POST /beers/1/reviews to be valid.")
//		t.FailNow()
//	}
//
//	newReviewMissing := true
//	for _, b := range db.FindReview(Review{BeerID: 1}) {
//		if b.BeerID == newReview.BeerID &&
//			b.FirstName == newReview.FirstName &&
//			b.LastName == newReview.LastName &&
//			b.Score == newReview.Score &&
//			b.Text == newReview.Text {
//			newReviewMissing = false
//		}
//	}
//
//	if newReviewMissing {
//		t.Errorf("Expected to find new entry in storage`")
//		t.FailNow()
//	}
//
//}
//
//func TestGetBeerReviews(t *testing.T) {
//	var reviewsFromRequest []Review
//	var reviewsFromStorage []Review
//
//	reviewsFromStorage, err := db.FindReviews(Review{BeerID: 1}
//	if err != nil {
//		t.Errorf("Failed to fetch reviews from storage: %s", err.Error())
//		t.FailNow()
//	}
//
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("GET", "/beers/1/reviews", nil)
//
//	router.ServeHTTP(w, r)
//
//	json.Unmarshal(w.Body.Bytes(), &reviewsFromRequest)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected route GET /beers/1/reviews to be valid.")
//		t.FailNow()
//	}
//
//	if len(reviewsFromRequest) != len(reviewsFromStorage) {
//		t.Error("Expected number of reviews from request to be the same as reviews in the storage")
//		t.FailNow()
//	}
//
//	var mapReviews = make(map[Review]int, len(reviewsFromStorage))
//	for _, review := range reviewsFromStorage {
//		mapReviews[review] = 1
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
