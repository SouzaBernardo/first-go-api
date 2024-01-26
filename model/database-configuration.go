package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	user := "postgres"
	password := "postgres"
	dbname := "postgres"
	host := "localhost"
	//host := "postgres-container"
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbname, password, host)
	open, err := sql.Open("postgres", connection)

	if err != nil {
		fmt.Println("CANNOT CONNECT IN DATABASE")
	}

	return open
}
