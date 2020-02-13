package utils

import (
	"github.com/google/uuid"
)

// NewUUID will generate a new uuid v4, panicing if the underlying source of
// randomness fails.
func NewUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return id.String()
}
