package main

import (
	"guilidb/src/handlers"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// setupConfigurations()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// app.Post("/get", getHandler)
	app.Post("/set", handlers.SetHandler)
	// app.Post("/audit", auditHandler)

	log.Fatal(app.Listen(":6644"))
}
