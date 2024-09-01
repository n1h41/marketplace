package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

func ConnectToDatabase() (err error) {
	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	Db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Panic(err)
		return
	}

	if err = Db.Ping(); err != nil {
		log.Panic(err)
		return
	}

	fmt.Println("Connected to database 🔥")
	return nil
}
