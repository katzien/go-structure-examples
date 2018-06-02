package main

import (
	"fmt"

	"github.com/katzien/structure-examples/domain-driven/storage"
	"github.com/katzien/structure-examples/domain-driven/adding"
	"github.com/katzien/structure-examples/domain-driven/reviewing"
)

func main() {

	// error handling omitted for simplicity
	beersStorage, _ := storage.NewJSONBeerStorage()
	reviewsStorage, _ := storage.NewJSONReviewStorage()

	// create the available services
	adder := adding.NewService(beersStorage)
	reviewer := reviewing.NewService(reviewsStorage)

	// add some sample data
	adder.AddSampleBeers()
	reviewer.AddSampleReviews()

	fmt.Println("Finished adding sample data.")
}
