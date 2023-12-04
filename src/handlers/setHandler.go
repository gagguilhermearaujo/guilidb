package handlers

import (
	"encoding/json"
	"guilidb/src/encrypting"
	"guilidb/src/filehandling"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetHandler(c *fiber.Ctx) error {
	documents, shouldReturnJsonErrorMessage, jsonErrorMessage := parseBody(c)
	if shouldReturnJsonErrorMessage {
		return jsonErrorMessage
	}

	shouldSkipRequest, responses := false, Responses{}
	for _, document := range documents.Documents {
		shouldSkipRequest, responses = validateRequest(responses, document)
		if shouldSkipRequest {
			continue
		}

		responses.Documents = setDocument(document, responses)
	}

	return c.JSON(responses)
}

func setDocument(document Request, responses Responses) []Response {
	document = addGuilidbFields(document)
	jsonDocument, _ := json.Marshal(document.Data)
	jsonDocument = append(jsonDocument, ',')
	encryptedDocument := encrypting.EncryptDocument(jsonDocument)
	filehandling.WriteFileToDisk(encryptedDocument, document.Collection, document.Key)

	return updateResponsesWithSetOkResult(document, responses)
}

func addGuilidbFields(document Request) Request {
	newData := document.Data
	newData["#guilidb.setAt"] = time.Now().UTC()
	newData["#guilidb.setBy"] = "xxx"

	return Request{
		Collection: document.Collection,
		Key:        document.Key,
		Data:       newData,
	}
}

func updateResponsesWithSetOkResult(document Request, responses Responses) []Response {
	response := Response{
		Collection: document.Collection,
		Key:        document.Key,
		StatusCode: fiber.StatusCreated,
		Message:    "Set command successful",
	}

	return append(responses.Documents, response)
}
