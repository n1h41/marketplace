package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	Db  *sql.DB
	Err error
)

func ConnectToDatabase() {
	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	Db, Err = sql.Open("postgres", connectionStr)
	if Err != nil {
		log.Fatalln(Err)
	}
	fmt.Println("Connected to database 🔥")
}
