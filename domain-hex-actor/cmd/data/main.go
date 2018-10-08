package main

import (
	"fmt"
	"time"

	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/reviewing"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/storage/json"
)

type Message interface{}

func main() {

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	// create the available services
	adder := adding.NewService(s)       // adding "actor"
	reviewer := reviewing.NewService(s) // reviewing "actor"

	resultsBeer := adder.AddSampleBeers(adding.DefaultBeers)
	resultsReview := reviewer.AddSampleReviews(reviewing.DefaultReviews)

	go func() {
		for result := range resultsBeer {
			fmt.Printf("Added sample beer with result %s.\n", result.GetMeaning()) // human-friendly
		}
	}()

	go func() {
		for result := range resultsReview {
			fmt.Printf("Added sample review with result %d.\n", result) // machine-friendly
		}
	}()

	// main could have its own "mailbox" exposed, for example an HTTP endpoint,
	// so we could be waiting here for more sample data to be added
	// (but we'll just exit for simplicity)

	time.Sleep(2 * time.Second) // this is here just to get the output from goroutines printed

	fmt.Println("No more data to add!")
}
