package utils

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/akfaew/test"
)

func TestError(t *testing.T) {
	myErr := fmt.Errorf("my error")

	t.Run("nil", func(t *testing.T) {
		test.True(t, Wrap(nil) == nil)
		test.True(t, Wrapf(nil, "something failed") == nil)
	})

	t.Run("errors.Is()", func(t *testing.T) {
		test.True(t, errors.Is(Wrap(myErr), myErr))
		test.True(t, errors.Is(Wrapf(myErr, "something failed"), myErr))
	})

	t.Run("paths", func(t *testing.T) {
		thisfile := "/utils/err_test.go:"
		test.True(t, strings.Contains(Wrap(myErr).Error(), thisfile))
		test.True(t, strings.Contains(Wrapf(myErr, "something failed").Error(), thisfile))
	})
}
