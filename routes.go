package main

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func getHandler(c *fiber.Ctx) error {
	bodyStruct, err := parseDocumentsFromBody(c)
	if err != nil {
		c.Status(422)
		return c.JSON(ErrorMessage{
			Error:   err.Error(),
			Example: map[string]any{"document": map[string]any{"firstName": "John", "lastName": "Doe"}},
		})
	}

	shouldSkip, responses := false, Responses{}
	for _, document := range bodyStruct.Documents {
		shouldSkip, responses = validateRequest(responses, document)
		if shouldSkip {
			continue
		}

		documentHistory, err := getDocumentHistoryFromDb(document.Collection, document.Key)
		var response DocumentResponse
		if err != nil {
			response = DocumentResponse{
				Collection: document.Collection,
				Key:        document.Key,
				StatusCode: fiber.StatusNotFound,
				Message:    err.Error(),
			}
		} else {
			response = DocumentResponse{
				Collection: document.Collection,
				Key:        document.Key,
				StatusCode: fiber.StatusOK,
				Message:    "Get command successful",
				Data:       documentHistory[len(documentHistory)-1],
			}
		}

		responses.Documents = append(responses.Documents, response)

	}
	return c.JSON(responses)
}

func setHandler(c *fiber.Ctx) error {
	bodyStruct, err := parseDocumentsFromBody(c)
	if err != nil {
		c.Status(422)
		return c.JSON(ErrorMessage{
			Error:   err.Error(),
			Example: map[string]any{"document": map[string]any{"firstName": "John", "lastName": "Doe"}},
		})
	}

	shouldSkip, responses := false, Responses{}
	for _, document := range bodyStruct.Documents {
		shouldSkip, responses = validateRequest(responses, document)
		if shouldSkip {
			continue
		}

		document.Data["#guilidb.setBy"] = "xxx"
		document.Data["#guilidb.setAt"] = time.Now().UTC()
		documentHistory, _ := getDocumentHistoryFromDb(document.Collection, document.Key)
		documentHistory = append(documentHistory, document.Data)

		documentToInsert, err := json.Marshal(documentHistory)
		handleError(err)

		encryptedDocument := encryptData(documentToInsert)
		writeDocumentFile(document.Collection, document.Key, encryptedDocument)
		response := DocumentResponse{
			Collection: document.Collection,
			Key:        document.Key,
			StatusCode: fiber.StatusCreated,
			Message:    "Set command successful",
		}
		responses.Documents = append(responses.Documents, response)
	}

	return c.JSON(responses)
}

func auditHandler(c *fiber.Ctx) error {
	bodyStruct, err := parseDocumentsFromBody(c)
	if err != nil {
		c.Status(422)
		return c.JSON(ErrorMessage{
			Error:   err.Error(),
			Example: map[string]any{"document": map[string]any{"firstName": "John", "lastName": "Doe"}},
		})
	}

	shouldSkip, responses := false, Responses{}
	for _, document := range bodyStruct.Documents {
		shouldSkip, responses = validateRequest(responses, document)
		if shouldSkip {
			continue
		}

		documentHistory, err := getDocumentHistoryFromDb(document.Collection, document.Key)
		var response DocumentResponse
		if err != nil {
			response = DocumentResponse{
				Collection: document.Collection,
				Key:        document.Key,
				StatusCode: fiber.StatusNotFound,
				Message:    err.Error(),
			}
		} else {
			response = DocumentResponse{
				Collection: document.Collection,
				Key:        document.Key,
				StatusCode: fiber.StatusOK,
				Message:    "Audit command successful",
				Audit:      documentHistory,
			}
		}

		responses.Documents = append(responses.Documents, response)

	}
	return c.JSON(responses)
}
