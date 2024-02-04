package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Imamsubekti26/Perpustakaan_Go/controllers"
	xlogger "github.com/Imamsubekti26/Perpustakaan_Go/utils/XLogger"
	"github.com/Imamsubekti26/Perpustakaan_Go/utils/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env variable
	err := godotenv.Load()
	if err != nil {
		xlogger.WriteAndShow(err)
	}

	// try to connect database
	db, err := database.Connection()
	if err != nil {
		xlogger.WriteAndShow(err)
	}

	// try to migrate database
	if err := db.Migrate(); err != nil {
		xlogger.WriteAndShow(err)
	}

	// create fiber instance
	app := fiber.New()

	// konfigurasi fiber
	app.Static("/assets", "./public/assets")

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userController := controllers.NewUserController(db.Connection)
	app.Post("/register", userController.Register)
	app.Post("/login", userController.Login)

	// running fiber server
	srv := fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	)
	log.Fatal(app.Listen(srv))
}
