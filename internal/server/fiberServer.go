package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"

	"n1h41/marketplace/internal/delivery/http/route"
	"n1h41/marketplace/internal/infrastructure/database"
)

type FiberServer struct{}

func NewFiberServer() Server {
	return &FiberServer{}
}

func (f *FiberServer) Run() {
	app := fiber.New()
	app.Use(favicon.New())
	app.Static("/static/", "./public")
	database.ConnectToDatabase()
	route.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
