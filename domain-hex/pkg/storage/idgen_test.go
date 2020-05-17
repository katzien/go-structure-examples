package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetID(t *testing.T) {
	id, err := GetID("testing")

	require.NoError(t, err)

	assert.True(t, strings.HasPrefix(id, "testing_"))
	assert.Len(t, id, 24)
}

func TestGetIDEmptyPrefix(t *testing.T) {
	id, err := GetID("")

	require.NoError(t, err)

	assert.True(t, strings.HasPrefix(id, "_"))
	assert.Len(t, id, 17)
}
