package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
)

func init() {

	conn, err := sql.Open("postgres", getDbConn())
	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(20)
	conn.SetMaxIdleConns(20)
	conn.SetConnMaxLifetime(time.Minute * 5)

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	db = conn
}

func getDbConn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}

// GetDBConnection returns a database connection to the application
func GetDBConnection() *sql.DB {
	return db
}
