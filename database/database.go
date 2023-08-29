package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// "os"
)

var db *sql.DB

func InitDB() {
	dataSourceName := "host=localhost port=5432 dbname=message_queue_db user=root password=postgres sslmode=disable"
	//  fmt.Sprintf(
	// 	"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
	// 	os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"), os.Getenv("DBUser"), os.Getenv("DBPass"),
	// )
	

	fmt.Println("dataSourceName:", dataSourceName)

	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("error:", err)
	}
}

func GetDB() *sql.DB {
	return db
}