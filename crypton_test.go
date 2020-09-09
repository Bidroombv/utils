package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var key = []byte{122, 37, 67, 42, 70, 45, 74, 97, 78, 100, 82, 103, 85, 107, 88, 112, 50, 114, 53, 117, 56, 120, 47, 65, 63, 68, 40, 71, 43, 75, 98, 80}

func TestDecrypt(t *testing.T) {
	t.Run("Decrypt", func(t *testing.T) {
		decrypted := Decrypt(iv, key, "pwinOs/QNg==")

		assert.Equal(t, "Foo Bar", decrypted)
	})
}

func TestEncrypt(t *testing.T) {
	t.Run("Encrypt", func(t *testing.T) {
		encrypted := Encrypt(iv, key, "Foo Bar")

		assert.Equal(t, "pwinOs/QNg==", encrypted)
	})
}
