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

type GetResponse struct {
	Document map[string]any `json:"document" xml:"document" form:"document"`
}

type AuditResponse struct {
	DocumentHistory []map[string]any `json:"documentHistory" xml:"documentHistory" form:"documentHistory"`
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
	return c.JSON(GetResponse{Document: document})
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

	documentHistory, _ := getDocumentHistoryFromDb(collection, key)
	documentHistory = append(documentHistory, bodyStruct.Document)

	documentToInsert, err := json.Marshal(documentHistory)
	handleError(err)

	encryptedDocument := encryptData(documentToInsert)
	writeDocumentFile(collection, key, encryptedDocument)

	return c.JSON(bodyStruct)
}

func auditHandler(c *fiber.Ctx) error {
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

	return c.JSON(AuditResponse{DocumentHistory: documentHistory})
}
