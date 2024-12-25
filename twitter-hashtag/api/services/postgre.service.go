package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	// Define the connection string
	connStr := "user=admin password=admin dbname=hashtag_count host=localhost port=5432 sslmode=disable"

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %v", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0)

	if err != nil {
		return nil, fmt.Errorf("error initializing the database: %v", err)
	}

	return db, nil
}
