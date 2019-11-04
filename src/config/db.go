package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbUser, dbPass)

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("failed to connect to db")
	}

	db.SetConnMaxLifetime(5 * time.Second)
}

func GetDB() *sql.DB {
	if db == nil {
		InitDB()
	}
	return db
}
