package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-json"
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

	_, err := os.Stat(fmt.Sprintf("%s/%s", dataDir, hashedCollection))
	if os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("%s/%s", dataDir, hashedCollection), 0700)
	}

	err = os.WriteFile(documentFileName, documentToInsert, 0600)
	handleError(err)
}

func getHashedCollectionAndFilename(collection string, key string) (uuid.UUID, string) {
	collectionNamespace := uuid.NewSHA1(dbConfig.NamespaceUUID, []byte(collection))
	hashedCollection := uuid.NewSHA1(collectionNamespace, []byte(collection))
	hashedKey := uuid.NewSHA1(collectionNamespace, []byte(key))
	documentFileName := fmt.Sprintf("%s/%s/%s.guilidb", dataDir, hashedCollection, hashedKey)
	return hashedCollection, documentFileName
}

// This function setup the config struct for DB and the data directory
func setupDirectoryAndConfigs() {
	// Create data directory
	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		os.Mkdir(dataDir, 0700)
	}

	encryptedKey, err := os.ReadFile(configPath)
	if err != nil {
		encryptedKey = setupNewDbConfigAndKey()
	}

	decryptedDocument := decryptData(encryptedKey)
	documentHistory := parseDocumentFileIntoDocumentHistory(decryptedDocument)
	latestConfig := documentHistory[len(documentHistory)-1]
	dbConfig = DbConfig{
		NamespaceUUID: uuid.MustParse(latestConfig["NamespaceUUID"].(string)),
	}

}

func setupNewDbConfigAndKey() []byte {
	newDbConfig := []DbConfig{
		{
			NamespaceUUID: uuid.New(),
		},
	}
	dbConfig = newDbConfig[0]
	newDbConfigBytes, _ := json.Marshal(newDbConfig)
	dbConfigToInsert := encryptData(newDbConfigBytes)
	_ = os.WriteFile(configPath, dbConfigToInsert, 0600)
	encryptedKey, _ := os.ReadFile(configPath)
	return encryptedKey
}
