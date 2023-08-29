package producer

import (
	"strconv"
	"context"
	"log"
	"github.com/segmentio/kafka-go"
	"os"
)

func SendMessage(productID int) {
	topic := os.Getenv("TOPIC_NAME")
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("BROKER_ADDRESS"), topic, partition)
	if err != nil {
		log.Fatal("Failed to connect to broker:", err)
	}
	defer conn.Close()

	message := kafka.Message{
		Value: []byte(strconv.Itoa(productID)),
	}

	_, err = conn.WriteMessages(message)
	if err != nil {
		log.Println("Failed to write message:", err)
	} else {
		log.Println("Message sent successfully")
	}
}