package config

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "tokopedia"
	dbname   = "postgres"
)

var (
	db *sql.DB
)

func Connect() {
	// Construct connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	d, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	// Test the connection
	err = d.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL")
	db = d
}

func GetDB() *sql.DB {
	return db
}
