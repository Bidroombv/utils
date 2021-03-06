package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var key = []byte{122, 37, 67, 42, 70, 45, 74, 97, 78, 100, 82, 103, 85, 107, 88, 112, 50, 114, 53, 117, 56, 120, 47, 65, 63, 68, 40, 71, 43, 75, 98, 80}

func TestDecrypt(t *testing.T) {
	t.Run("Decrypt", func(t *testing.T) {
		tcs := []*struct {
			encrypted string
			plain     string
		}{
			{
				"pwinOs/QNg==",
				"Foo Bar",
			},
			{
				"hwinWujJJcS7nIbEySJq",
				"foo@example.com",
			},
			{
				"ylPwOryDd4n/xdXKnXU+",
				"+48 123 456 789",
			},
			{
				"hwinN+/QNg==",
				"foo-bar",
			},
			{
				"wCfrPqjvYoPj2bzB",
				"!@#$%^&*()_+",
			},
		}

		var decrypted string
		for _, tc := range tcs {
			decrypted, _ = Decrypt(iv, key, tc.encrypted)
			assert.Equal(t, tc.plain, decrypted)
		}
	})
}

func TestEncrypt(t *testing.T) {
	t.Run("Encrypt", func(t *testing.T) {
		tcs := []*struct {
			plain     string
			encrypted string
		}{
			{
				"Foo Bar",
				"pwinOs/QNg==",
			},
			{
				"foo@example.com",
				"hwinWujJJcS7nIbEySJq",
			},
			{
				"+48 123 456 789",
				"ylPwOryDd4n/xdXKnXU+",
			},
			{
				"foo-bar",
				"hwinN+/QNg==",
			},
			{
				"!@#$%^&*()_+",
				"wCfrPqjvYoPj2bzB",
			},
		}

		var encrypted string
		for _, tc := range tcs {
			encrypted, _ = Encrypt(iv, key, tc.plain)
			assert.Equal(t, tc.encrypted, encrypted)
		}
	})
}
