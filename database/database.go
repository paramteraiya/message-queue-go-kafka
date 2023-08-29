package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func InitDB() {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
	)

	fmt.Println("dataSourceName:", dataSourceName)

	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("error:", err)
	}

	createTablesIfNotExist()
}

func GetDB() *sql.DB {
	return db
}

func createTablesIfNotExist() {
	
	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name varchar(255),
			mobile varchar(255),
			latitude double precision,
			longitude double precision,
			created_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp
		)
	`

	_, err := db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal("error creating table:", err)
	}

	createProductsTableQuery := `
		CREATE TABLE IF NOT EXISTS products (
			product_id serial PRIMARY KEY,
			product_name varchar(255),
			product_description text,
			product_images text[],
			product_price double precision,
			compressed_product_images text[],
			created_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp
		)
	`

	_, err1 := db.Exec(createProductsTableQuery)
	if err1 != nil {
		log.Fatal("error creating table:", err1)
	}
	
}




