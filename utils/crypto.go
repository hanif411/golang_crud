package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func Encrypt(key []byte, text string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "error buat aes"
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "error buat cipher new gcm"
	}

	nonce := make([]byte, aesGCM.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := aesGCM.Seal(nonce, nonce, []byte(text), nil)

	return hex.EncodeToString(cipherText)
}

func Decrypt(key []byte, cipherText string) string {
	data, err := hex.DecodeString(cipherText)
	if err != nil {
		return "error decode chiper text"
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "error buat aes"
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "error aes gcm"
	}

	nonSize := aesGCM.NonceSize()
	nonce, purechipertext := data[:nonSize], data[nonSize:]

	text, err := aesGCM.Open(nil, nonce, purechipertext, nil)
	if err != nil {
		return "error open gcm decrypt"
	}

	return string(text)
}
