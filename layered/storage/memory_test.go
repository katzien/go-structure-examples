package storage

import (
	"testing"
	"time"
	"github.com/katzien/structure-examples/layered/models"
)

func TestSaveBeer(t *testing.T) {

	storage := new(Memory)
	sampleBeer := models.Beer{
		ID:      1,
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
		Created: time.Date(2017, time.October, 24, 22, 6, 0, 0, time.UTC),
	}
	storage.SaveBeer(sampleBeer)

	if len(storage.cellar) == 0 {
		t.Errorf("Expected sample beer to be added to storage list.")
		t.FailNow()
	}

	sampleBeerChange := sampleBeer
	sampleBeerChange.Name = "Not a beer name"
	storage.SaveBeer(sampleBeerChange)

	if len(storage.cellar) > 1 {
		t.Errorf("Expected sample beer to be updated instead of creating another entry.")
		t.FailNow()
	}

	if storage.cellar[0].Name == sampleBeer.Name {
		t.Errorf("Expected sample beer name to be updated with new name.")
		t.FailNow()
	}
}
