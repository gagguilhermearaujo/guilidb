package encrypting

import (
	"crypto/rand"
	rnd "math/rand"

	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/chacha20poly1305"
)

var NAMESPACEUUID = uuid.New()
var ADMINPASSWORD, _ = password.Generate(32, rnd.Intn(4)+4, rnd.Intn(4)+4, false, true)
var AEAD, _ = chacha20poly1305.NewX([]byte(ADMINPASSWORD))

func EncryptDocument(document []byte) []byte {
	nonce := createRandomNonce(document)
	encryptedDocument := AEAD.Seal(nonce, nonce, document, nil)

	return encryptedDocument
}

func createRandomNonce(dataToEncrypt []byte) []byte {
	nonce := make([]byte, AEAD.NonceSize(), AEAD.NonceSize()+len(dataToEncrypt)+AEAD.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	return nonce
}

func GetHashedCollectionAndKey(collection string, key string) (string, string) {
	hashedCollection := uuid.NewSHA1(NAMESPACEUUID, []byte(collection)).String()
	hashedKey := uuid.NewSHA1(NAMESPACEUUID, []byte(key)).String()

	return hashedCollection, hashedKey
}
