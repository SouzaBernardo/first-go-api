package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func ConnectDB() *sql.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbname, password, host)
	open, err := sql.Open("postgres", connection)

	if err != nil {
		fmt.Println("CANNOT CONNECT IN DATABASE")
	}

	return open
}
