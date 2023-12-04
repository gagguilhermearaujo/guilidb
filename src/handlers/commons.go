package handlers

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

var (
	hasInvalidKey, _ = regexp.Compile(`[^A-Za-z0-9.\-_]+`)
)

func parseBody(c *fiber.Ctx) (Requests, bool, error) {
	var documents Requests
	err := c.BodyParser(&documents)
	if err != nil {
		c.Status(422)
		return Requests{}, true, c.JSON(ErrorMessage{
			Error:   err.Error(),
			Example: map[string]any{"document": map[string]any{"firstName": "John", "lastName": "Doe"}},
		})
	}
	return documents, false, nil
}

func validateRequest(responses Responses, document Request) (bool, Responses) {
	validRequest := true
	err := validateCollectionAndKey(document.Collection, document.Key)
	if err != nil {
		response := Response{
			Collection: document.Collection,
			Key:        document.Key,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
		}
		responses.Documents = append(responses.Documents, response)
		validRequest = false
	}

	return validRequest, responses
}

func validateCollectionAndKey(collection string, key string) error {
	if hasInvalidKey.MatchString(collection) {
		return fmt.Errorf("wrong collection name '%s', it should be used only letters, numbers, dots, underscores or dashes", collection)
	}
	if hasInvalidKey.MatchString(key) {
		return fmt.Errorf("wrong key name '%s', it should be used only letters, numbers, dots, underscores or dashes", key)
	}
	return nil
}
