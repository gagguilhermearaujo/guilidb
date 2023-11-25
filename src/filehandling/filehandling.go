package filehandling

import (
	"fmt"
	"guilidb/src/encrypting"
	"os"

	"github.com/google/uuid"
)

const DATADIR = "data"

func WriteFileToDisk(documentToInsert []byte, collection string, key string) {
	hashedCollection, hashedKey := encrypting.GetHashedCollectionAndKey(collection, key)
	documentParentPath := fmt.Sprintf("%s/%s/%s", DATADIR, hashedCollection, hashedKey)
	documentFullPath := fmt.Sprintf("%s/%s.guilidb", documentParentPath, uuid.NewString())

	os.MkdirAll(documentParentPath, 0700)
	os.WriteFile(documentFullPath, documentToInsert, 0600)

	return
}
