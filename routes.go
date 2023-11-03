package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Error   string         `json:"error"`
	Example map[string]any `json:"example,omitempty"`
}

type BodyStruct struct {
	Document map[string]any `json:"document" xml:"document" form:"document"`
}

func getHandler(c *fiber.Ctx) error {
	collection, key, err := validateCollectionAndKey(c)
	if err != nil {
		c.Status(400)
		return c.JSON(ErrorMessage{Error: err.Error()})
	}

	documentHistory, err := getDocumentHistoryFromDb(collection, key)
	if err != nil {
		c.Status(404)
		return c.JSON(ErrorMessage{Error: err.Error()})
	}

	document := documentHistory[len(documentHistory)-1]
	return c.JSON(BodyStruct{Document: document})
}

func setHandler(c *fiber.Ctx) error {
	collection, key, err := validateCollectionAndKey(c)
	if err != nil {
		c.Status(400)
		return c.JSON(ErrorMessage{Error: err.Error()})
	}

	bodyStruct, err := parseDocumentFromBody(c)
	if err != nil {
		c.Status(422)
		return c.JSON(ErrorMessage{
			Error:   err.Error(),
			Example: map[string]any{"document": map[string]any{"firstName": "John", "lastName": "Doe"}},
		})
	}

	documentHistory, err := getDocumentHistoryFromDb(collection, key)
	if err != nil {
		c.Status(404)
		return c.JSON(ErrorMessage{Error: err.Error()})
	}

	documentToInsert, err := json.Marshal(bodyStruct.Document)
	handleError(err)

	encryptedDocument := encryptData(documentToInsert)
	writeDocumentFile(collection, key, encryptedDocument)

	return c.JSON(bodyStruct)
}
