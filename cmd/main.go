package main

import (
	"log"

	"github.com/joho/godotenv"

	"n1h41/marketplace/internal/server"
)

func main() {
	// INFO: Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	fiberServer := server.NewFiberServer()
	fiberServer.Run()
}
