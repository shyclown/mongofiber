package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbname   = "fiber_local_01"
)

func Connect() {
	var err error
	// Use DSN string to open
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			user, password, host, port, dbname,
		))

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db

	fmt.Println("😁 Connected to database")
}
