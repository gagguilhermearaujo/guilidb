package main

import (
	"crypto/rand"

	"golang.org/x/crypto/chacha20poly1305"
)

var (
	key     = make([]byte, chacha20poly1305.KeySize)
	aead, _ = chacha20poly1305.NewX(key)
)

func encryptData(dataToEncrypt []byte) []byte {
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(dataToEncrypt)+aead.Overhead())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	encryptedData := aead.Seal(nonce, nonce, dataToEncrypt, nil)
	return encryptedData
}

func decryptData(ciphertext []byte) []byte {
	nonce, ciphertext := ciphertext[:aead.NonceSize()], ciphertext[aead.NonceSize():]
	decryptedData, _ := aead.Open(nil, nonce, ciphertext, nil)

	return decryptedData
}
