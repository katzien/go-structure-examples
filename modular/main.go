package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/structure-examples/modular/database"
	"github.com/katzien/structure-examples/modular/beers"
	"github.com/katzien/structure-examples/modular/reviews"
)

var router *httprouter.Router

func init() {
	var err error

	err = database.NewStorage(database.Memory)
	if err != nil {
		log.Fatal(err)
	}

	database.PopulateBeers()
	database.PopulateReviews()

	router = httprouter.New()

	router.GET("/beers", beers.GetBeers)
	router.GET("/beers/:id", beers.GetBeer)
	router.GET("/beers/:id/reviews", reviews.GetBeerReviews)

	router.POST("/beers", beers.AddBeer)
	router.POST("/beers/:id/reviews", reviews.AddBeerReview)
}

func main() {
	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
