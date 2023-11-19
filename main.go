package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func setupDirectoryAndConfigs() {
	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		os.Mkdir(dataDir, 0700)
	}

	dbConfigs, err := getDocumentHistoryFromDb(configCollection, configKey)
	if err != nil {
		newDbConfig := []DbConfig{{NamespaceUUID: uuid.New()}}
		newDbConfigBytes, _ := json.Marshal(newDbConfig)
		dbConfigToInsert := encryptData(newDbConfigBytes)
		writeDocumentFile(configCollection, configKey, dbConfigToInsert)
		dbConfig = newDbConfig[0]
	} else {
		dbConfigsStruct, _ := json.Marshal(dbConfigs[len(dbConfigs)-1])
		json.Unmarshal(dbConfigsStruct, &dbConfig)
	}
}
