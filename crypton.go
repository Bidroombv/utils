package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(iv []byte, key []byte, plain string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plain)
	ctr := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plainBytes))
	ctr.XORKeyStream(ciphertext, plainBytes)

	return encodeBase64(ciphertext), nil
}

func Decrypt(iv []byte, key []byte, encrypted string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := decodeBase64(encrypted)
	ctr := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	ctr.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
