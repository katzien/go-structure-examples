package main

import "time"

// Review defines the properties of a beer review
type Review struct {
	ID        string
	BeerID    int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     int    `json:"score"`
	Text      string `json:"text"`
	Created   time.Time
}

// Beer defines the properties of a beer
type Beer struct {
	ID        int
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Abv       float32 `json:"abv"`
	ShortDesc string  `json:"short_description"`
	Created   time.Time
}
