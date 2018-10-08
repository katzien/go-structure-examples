package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingMemoryStorage(t *testing.T) {
	s, err := NewStorage(Memory)

	assert.Nil(t, err)
	assert.IsType(t, &StorageMemory{}, s)
}

func TestCreatingJSONStorage(t *testing.T) {
	s, err := NewStorage(JSON)

	assert.Nil(t, err)
	assert.IsType(t, &StorageJSON{}, s)
}
