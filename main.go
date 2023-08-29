package main

import (
	"github.com/message-queue/api"
	"github.com/message-queue/database"
	"github.com/message-queue/consumer"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// Load the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	// initialize the db
	database.InitDB()

	// start api server
	go api.StartServer()
	
	// start consumer
	consumer.StartConsumer()

}