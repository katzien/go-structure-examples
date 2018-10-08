package json

import "time"

// Beer defines the storage form of a beer
type Beer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}
