package main

import (
	"log"

	"github.com/Imamsubekti26/Perpustakaan_Go/utils"
	xlogger "github.com/Imamsubekti26/Perpustakaan_Go/utils/XLogger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnv()

	app := fiber.New()

	_, err := utils.InitDB()
	if err != nil {
		xlogger.WriteAndShow(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
