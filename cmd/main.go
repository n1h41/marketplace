package main

import (
	"log"

	"github.com/joho/godotenv"
	"n1h41/marketplace/internal/server"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	fiberServer := server.NewFiberServer()
	fiberServer.Run()
}

