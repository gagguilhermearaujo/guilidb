package main

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func parseDocumentFromBody(c *fiber.Ctx) (*BodyStruct, error) {
	document := new(BodyStruct)
	err := c.BodyParser(document)
	if err != nil {
		return nil, fmt.Errorf("body is not a valid JSON, XML or Form. It should contain a document/object/dict/hashmap at 'document' key")
	}
	return document, nil
}

func parseDocumentFileIntoDocumentHistory(documentFile []byte) []map[string]interface{} {
	var documentHistory []map[string]interface{}
	err := json.Unmarshal(documentFile, &documentHistory)
	handleError(err)
	return documentHistory
}
