package db

import (
	"database/sql"
	"log"
	"time"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"

)

var DB *sql.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	var err error

	for i := 0; i < 10; i++ {

		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Println("Waiting for database...")
			time.Sleep(2 * time.Second)
			continue
		}

		err = DB.Ping()
		if err == nil {
			log.Println("Successfully connected to database")
			return
		}

		log.Println("Database not ready yet, retrying...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Could not connect to database")
}