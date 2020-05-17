package memory

import (
	"time"
)

// Beer defines the storage form of a beer
type Beer struct {
	ID        string
	Name      string
	Brewery   string
	Abv       float32
	ShortDesc string
	Created   time.Time
}
