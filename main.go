package main

import (
	"log"

	"github.com/Imamsubekti26/Perpustakaan_Go/utils"
	xlogger "github.com/Imamsubekti26/Perpustakaan_Go/utils/XLogger"
	"github.com/Imamsubekti26/Perpustakaan_Go/utils/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnv()

	app := fiber.New()

	db, err := database.Connection()
	if err != nil {
		xlogger.WriteAndShow(err)
	}

	db.Migrate()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
