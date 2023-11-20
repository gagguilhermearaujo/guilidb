package main

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	setupDirectoryAndConfigs()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Use /get/{key} for getting the list of documents")
	})

	app.Post("/get", getHandler)
	app.Post("/set", setHandler)
	app.Post("/audit", auditHandler)

	log.Fatal(app.Listen(":6644"))
}
