package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func ConnectToDatabase() {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to database 🔥")
}
