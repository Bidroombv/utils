package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	id := NewUUID()
	assert.Len(t, id, 36)
}
