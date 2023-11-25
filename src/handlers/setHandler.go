package handlers

import (
	"encoding/json"
	"guilidb/src/encrypting"
	"guilidb/src/filehandling"

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
	jsonDocument, _ := json.Marshal(document.Data)
	encryptedDocument := encrypting.EncryptDocument(jsonDocument)
	filehandling.WriteFileToDisk(encryptedDocument, document.Collection, document.Key)

	return updateResponsesWithSetOkResult(document, responses)
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
