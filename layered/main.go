package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/go-structure-examples/layered/handlers"
	"github.com/katzien/go-structure-examples/layered/storage"
)

var router *httprouter.Router

func init() {
	var err error

	err = storage.NewStorage(storage.MemoryType)
	if err != nil {
		log.Fatal(err)
	}

	PopulateBeers()
	PopulateReviews()

	router = httprouter.New()

	router.GET("/beers", handlers.GetBeers)
	router.GET("/beers/:id", handlers.GetBeer)
	router.GET("/beers/:id/reviews", handlers.GetBeerReviews)

	router.POST("/beers", handlers.AddBeer)
	router.POST("/beers/:id/reviews", handlers.AddBeerReview)
}

func main() {
	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
