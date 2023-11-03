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
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Use /get/{key} for getting the list of documents")
	})

	app.Get("/get/:collection/:key", getHandler)
	app.Post("/set/:collection/:key", setHandler)

	log.Fatal(app.Listen(":6644"))
}
