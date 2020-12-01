package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	fmt.Println("Connecting to Database...")

	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/tinapp")

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Database connected")

	return db
}
