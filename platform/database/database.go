package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func ConnectToDatabase() {
	connectionStr := "user=n1h41 password=nihal@23ktu dbname=marketplace sslmode=disable"
	db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to database 🔥")
}
