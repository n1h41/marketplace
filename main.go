package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/joho/godotenv"

	"n1h41/marketplace/platform/database"
)

func main() {
	// INFO: Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(favicon.New())

	app.Static("/static", "./public/")

	database.ConnectToDatabase()

	Setup(app)

	log.Fatal(app.Listen(":3000"))
}
