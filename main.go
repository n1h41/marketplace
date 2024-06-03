package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"n1h41/marketplace/handlers"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./public/")

	handlers.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
