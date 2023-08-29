package main

import (
	"github.com/message-queue/api"
	"github.com/message-queue/database"
	"github.com/message-queue/consumer"
)

func main() {
	// initialize the db
	database.InitDB()

	// start api server
	go api.StartServer()
	
	// start consumer
	consumer.StartConsumer()

}