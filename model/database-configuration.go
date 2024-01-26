package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	user := "postgres"
	password := "postgres"
	dbname := "postgres"
	host := "postgres-container"
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbname, password, host)

	return sql.Open("postgres", connection)
}
