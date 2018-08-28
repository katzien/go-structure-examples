package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ServerAddr defines the http host and port of the beer server
const ServerAddr = "localhost:8080"

var db Storage
var router *httprouter.Router

// note: avoid using init
func init() {
	var err error

	db, err = NewStorage(JSON)
	if err != nil {
		log.Fatal(err)
	}

	PopulateBeers()
	PopulateReviews()

	router = httprouter.New()

	router.GET("/beers", GetBeers)
	router.GET("/beers/:id", GetBeer)
	router.GET("/beers/:id/reviews", GetBeerReviews)

	router.POST("/beers", AddBeer)
	router.POST("/beers/:id/reviews", AddBeerReview)
}

func main() {
	fmt.Println("The beer server is on tap at http://localhost:8080.")
	log.Fatal(http.ListenAndServe(ServerAddr, router))
}
