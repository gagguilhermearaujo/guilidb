package main

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

var (
	hasInvalidKey, _ = regexp.Compile(`[^A-Za-z0-9.\-_]+`)
)

func validateRequest(responses Responses, document DocumentRequest) (bool, Responses) {
	shouldSkip := false
	err := validateCollectionAndKey(document.Collection, document.Key)
	if err != nil {
		response := DocumentResponse{
			Collection: document.Collection,
			Key:        document.Key,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
		}
		responses.Documents = append(responses.Documents, response)
		shouldSkip = true
	}
	return shouldSkip, responses
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
