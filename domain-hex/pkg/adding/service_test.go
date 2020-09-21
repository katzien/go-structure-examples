package adding

import (
	"github.com/katzien/go-structure-examples/domain-hex/pkg/listing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddBeers(t *testing.T) {
	b1 := Beer{
		Name:      "Test Beer 1",
		Brewery:   "Brewery One",
		Abv:       3.6,
		ShortDesc: "Lorem Ipsum",
	}

	b2 := Beer{
		Name:      "Test Beer 2",
		Brewery:   "Brewery Two",
		Abv:       4.8,
		ShortDesc: "Bacon Ipsum",
	}

	mR := new(mockStorage)

	s := NewService(mR)

	err := s.AddBeer(b1, b2)
	require.NoError(t, err)

	beers := mR.GetAllBeers()
	assert.Len(t, beers, 2)
}

type mockStorage struct {
	beers []Beer
}

func (m *mockStorage) AddBeer(b Beer) error {
	m.beers = append(m.beers, b)

	return nil
}

func (m *mockStorage) GetAllBeers() []listing.Beer {
	beers := []listing.Beer{}

	for _, bb := range m.beers {
		b := listing.Beer{
			Name:      bb.Name,
			Brewery:   bb.Brewery,
			Abv:       bb.Abv,
			ShortDesc: bb.ShortDesc,
		}
		beers = append(beers, b)
	}

	return beers
}
