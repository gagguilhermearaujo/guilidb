package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

func getDocumentHistoryFromDb(collection string, key string) ([]map[string]any, error) {
	_, documentFileName := getHashedCollectionAndFilename(collection, key)

	encryptedDocument, err := os.ReadFile(documentFileName)
	if err != nil {
		return []map[string]any{}, fmt.Errorf("error: could not find key '%s' on collection '%s'", key, collection)
	}

	decryptedDocument := decryptData(encryptedDocument)
	documentHistory := parseDocumentFileIntoDocumentHistory(decryptedDocument)

	return documentHistory, nil
}

func writeDocumentFile(collection string, key string, documentToInsert []byte) {
	hashedCollection, documentFileName := getHashedCollectionAndFilename(collection, key)

	_, err := os.Stat(fmt.Sprintf("data/%s", hashedCollection))
	if os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("data/%s", hashedCollection), 0700)
	}

	err = os.WriteFile(documentFileName, documentToInsert, 0600)
	handleError(err)
}

func getHashedCollectionAndFilename(collection string, key string) (uuid.UUID, string) {
	namespace := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(collection))
	hashedCollection := uuid.NewSHA1(namespace, []byte(collection))
	hashedKey := uuid.NewSHA1(namespace, []byte(key))
	documentFileName := fmt.Sprintf("data/%s/%s.guilidb", hashedCollection, hashedKey)
	return hashedCollection, documentFileName
}
