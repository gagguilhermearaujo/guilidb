package filehandling

import (
	"fmt"
	"guilidb/src/encrypting"
	"os"
)

const DATADIR = "data"

func WriteFileToDisk(documentToInsert []byte, collection string, key string) {
	hashedCollection, hashedKey := encrypting.GetHashedCollectionAndKey(collection, key)
	collectionPath := fmt.Sprintf("%s/%s", DATADIR, hashedCollection)
	documentFullPath := fmt.Sprintf("%s/%s.guilidb", collectionPath, hashedKey)

	os.MkdirAll(collectionPath, 0700)
	f, err := os.OpenFile(documentFullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write(documentToInsert)
	if err != nil {
		panic(err)
	}

	return
}

func GetFileFromDisk(collection string, key string) []byte {
	hashedCollection, hashedKey := encrypting.GetHashedCollectionAndKey(collection, key)
	documentFullPath := fmt.Sprintf("%s/%s/%s.guilidb", DATADIR, hashedCollection, hashedKey)

	encryptedDocument, err := os.ReadFile(documentFullPath)
	if err != nil {
		panic(err)
	}

	return encryptedDocument

}
