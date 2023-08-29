package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"github.com/message-queue/models"
	"log"
	"path/filepath"
	"net/http"
	"github.com/google/uuid"
	"image"
	"github.com/disintegration/imaging"
	"strings"
)

func StartConsumer() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("BROKER_ADDRESS")},
		Topic:     os.Getenv("TOPIC_NAME"),
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
		GroupID:   "my-group",
		StartOffset: kafka.LastOffset,
	})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	fmt.Println("Waiting for messages...")

	for {
		select {
		case <-signals:
			fmt.Println("Interrupt signal received")
			r.Close()
			return
		default:
			msg, err := r.ReadMessage(context.Background())
			if err == nil {
				fmt.Printf("Received message: %s\n", string(msg.Value))
				r.CommitMessages(context.Background(), msg)
				productID, err := strconv.Atoi(string(msg.Value))
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				result, err := downloadAndCompressImages(productID)
				fmt.Println("result: ", result)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			} else {
				fmt.Println("Error reading message:", err)
			}
		}
	}
}


func downloadAndCompressImages(productID int) ([]string, error) {
	imageURLs := models.GetImages(productID)
	compressedPaths := make([]string, 0, len(imageURLs))

	for _, imageURL := range imageURLs {
		imageURL = strings.Trim(imageURL, "{}")
		urls := strings.Split(imageURL, ",") 
		for _, url := range urls {
			url = strings.TrimSpace(url) 
			fmt.Println("url", url)
			
			response, err := http.Get(url)
			if err != nil {
				log.Printf("Failed to download image from %s: %v\n", url, err)
				continue
			}
			defer response.Body.Close()

			_, dir_err := os.Stat("compressed_images")
			if os.IsNotExist(dir_err) {
				// Directory doesn't exist, create it
				err := os.MkdirAll("compressed_images", os.ModePerm)
				if err != nil {
					fmt.Println("Error creating directory:", err)
				}
				fmt.Println("Directory created:", "compressed_images")
			} else if dir_err != nil {
				fmt.Println("Error:", dir_err)
			}

			// Create a unique filename for the compressed image
			compressedFileName := fmt.Sprintf("compressed_%s.jpg", uuid.New().String())
			compressedPath := filepath.Join("compressed_images", compressedFileName)

			// Decode the downloaded image
			img, _, err := image.Decode(response.Body)
			if err != nil {
				log.Printf("Failed to decode image from %s: %v\n", url, err)
				continue
			}

			// Resize and compress the image
			dst := imaging.Resize(img, 800, 600, imaging.Lanczos)
			err = imaging.Save(dst, compressedPath)
			if err != nil {
				log.Printf("Failed to save compressed image: %v\n", err)
				continue
			}

			compressedPaths = append(compressedPaths, compressedPath)
			models.StoreCompressedImages(productID, compressedPaths)
		}
	}

	return compressedPaths, nil
}




