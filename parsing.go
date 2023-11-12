package main

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func parseDocumentsFromBody(c *fiber.Ctx) (*Requests, error) {
	documents := new(Requests)
	err := c.BodyParser(documents)
	if err != nil {
		return nil, fmt.Errorf("body is not a valid JSON, XML or Form. It should contain a document/object/dict/hashmap at 'documents' key")
	}
	return documents, nil
}

func parseDocumentFileIntoDocumentHistory(documentFile []byte) []map[string]interface{} {
	var documentHistory []map[string]interface{}
	err := json.Unmarshal(documentFile, &documentHistory)
	handleError(err)
	return documentHistory
}
