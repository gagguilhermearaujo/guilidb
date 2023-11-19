package main

import (
	"crypto/rand"
)

func encryptData(dataToEncrypt []byte) []byte {
	_, err := rand.Read(encryptionKey)
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
