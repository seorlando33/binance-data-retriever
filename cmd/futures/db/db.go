package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDBConnection() *sql.DB {

	once.Do(func() {
		conn, err := sql.Open("postgres", GetPgConn())
		if err != nil {
			panic(err.Error())
		}

		conn.SetMaxOpenConns(20)
		conn.SetMaxIdleConns(20)
		conn.SetConnMaxLifetime(time.Minute * 5)

		err = conn.Ping()
		if err != nil {
			panic(err)
		}

		db = conn
	})

	return db
}

func GetPgConn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}