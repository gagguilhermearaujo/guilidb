package main

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

var (
	hasInvalidKey, _ = regexp.Compile(`[^A-Za-z0-9.\-_]+`)
)

func validateCollectionAndKey(c *fiber.Ctx) (string, string, error) {
	for _, validatedKey := range []string{"collection", "key"} {
		keyValue := c.Params(validatedKey)
		if hasInvalidKey.MatchString(keyValue) {
			c.SendStatus(400)
			return "", "", fmt.Errorf("wrong %s name '%s', it should be used only letters, numbers, dots, underscores or dashes", validatedKey, keyValue)
		}
	}
	return c.Params("collection"), c.Params("key"), nil
}
