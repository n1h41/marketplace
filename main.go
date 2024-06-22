package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"n1h41/marketplace/handlers"
	"n1h41/marketplace/platform/database"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./public/")

	handlers.Setup(app)

	database.ConnectToDatabase()

	log.Fatal(app.Listen(":3000"))
}
