package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/structure-examples/domain-driven/beers"
	"github.com/katzien/structure-examples/domain-driven/reviews"
	"github.com/katzien/structure-examples/domain-driven/storage"
	"github.com/katzien/structure-examples/domain-driven/adding"
	"github.com/katzien/structure-examples/domain-driven/reviewing"
	"github.com/katzien/structure-examples/domain-driven/listing"
)

func main() {

	// set up storage
	storageType := storage.InMemory // this could be a flag; hardcoded here for simplicity

	var beersStorage beers.Repository
	var reviewsStorage reviews.Repository

	switch storageType {
	case storage.InMemory:
		beersStorage = new(storage.MemoryBeerStorage)
		reviewsStorage = new(storage.MemoryReviewStorage)
	case storage.JSONFiles:
		// error handling omitted for simplicity
		beersStorage, _ = storage.NewJSONBeerStorage();
		reviewsStorage, _ = storage.NewJSONReviewStorage();
	}

	// create the available services
	adder := adding.NewService(beersStorage)
	reviewer := reviewing.NewService(reviewsStorage)
	lister := listing.NewService(beersStorage, reviewsStorage)

	// add some sample data
	adder.AddSampleBeers()
	reviewer.AddSampleReviews()

	// set up the HTTP server
	router := httprouter.New()

	router.GET("/beers", listing.MakeGetBeersEndpoint(lister))
	router.GET("/beers/:id", listing.MakeGetBeerEndpoint(lister))
	router.GET("/beers/:id/reviews", listing.MakeGetBeerReviewsEndpoint(lister))

	router.POST("/beers", adding.MakeAddBeerEndpoint(adder))
	router.POST("/beers/:id/reviews", reviewing.MakeAddBeerReviewEndpoint(reviewer))

	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
