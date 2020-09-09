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

func Encrypt(iv []byte, key []byte, plain string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plainBytes := []byte(plain)
	cfb := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plainBytes))
	cfb.XORKeyStream(ciphertext, plainBytes)

	return encodeBase64(ciphertext)
}

func Decrypt(iv []byte, key []byte, encrypted string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := decodeBase64(encrypted)
	cfb := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)

	return string(plaintext)
}
