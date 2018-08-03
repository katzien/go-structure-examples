package main

import (
	"github.com/katzien/go-structure-examples/actor/pkg/adding"
	"github.com/katzien/go-structure-examples/actor/pkg/reviewing"
	"fmt"
	"github.com/katzien/go-structure-examples/actor/storage"
	"github.com/katzien/go-structure-examples/actor/pkg/beers"
	"github.com/katzien/go-structure-examples/actor/pkg/reviews"
)

type Message interface{}

func main() {

	// error handling omitted for simplicity
	beersStorage, _ := storage.NewJSONBeerStorage()
	reviewsStorage, _ := storage.NewJSONReviewStorage()

	// create the available services
	adder := adding.NewService(beersStorage)
	reviewer := reviewing.NewService(reviewsStorage)

	beersDone := make(chan adding.Event)
	reviewsDone := make(chan reviewing.Event)

	go adder.AddSampleBeers(beersDone, beers.DefaultBeers)
	go reviewer.AddSampleReviews(reviewsDone, reviews.DefaultReviews)

	resultBeers := <-beersDone
	fmt.Printf("Finished adding sample beers with result %s.\n", resultBeers.GetMeaning()) // human-friendly

	resultReviews := <-reviewsDone
	fmt.Printf("Finished adding sample reviews with result %d.\n", resultReviews) // machine-friendly
}
