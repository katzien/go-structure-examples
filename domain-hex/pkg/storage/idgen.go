package storage

import (
	"crypto/rand"
	"fmt"
)

// GetID returns a random ID string of the format prefix_random16chars, e.g. beer_ts72nf6ak8dts73g.
// This is a simple (naive) implementation using the rand package,
// just to avoid importing external UUID packages in this demo app.
// This implementation in no way guarantees uniqueness, so please don't use it for any production purposes!
func GetID(prefix string) (string, error) {
 	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%x", prefix, b), nil
}
