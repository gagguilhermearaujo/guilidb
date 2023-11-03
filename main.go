package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0777)
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Use /get/{key} for getting the list of documents")
	})

	app.Get("/get/:collection/:key", getHandler)
	app.Post("/set/:collection/:key", setHandler)
	app.Get("/audit/:collection/:key", auditHandler)

	log.Fatal(app.Listen(":6644"))
}
