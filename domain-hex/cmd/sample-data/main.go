package main

import (
	"fmt"
	"github.com/katzien/go-structure-examples/domain-hex/storage/json"
	"github.com/katzien/go-structure-examples/domain-hex/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex/pkg/reviewing"
)

func main() {

	var adder adding.Service
	var reviewer reviewing.Service

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)
	reviewer = reviewing.NewService(s)

	// add some sample data
	adder.AddSampleBeers()
	reviewer.AddSampleReviews()

	fmt.Println("Finished adding sample data.")
}
