package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/message-queue/database"
	"github.com/message-queue/models"
	"github.com/message-queue/producer"
)

type PostData struct {
	UserID             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	ProductPrice       float64  `json:"product_price"`
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var data PostData
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Println("err: ", err)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println(data)
	
	// Create Product object
	product := models.Product{
		// UserID:      userID,
		ProductName: data.ProductName,
		ProductDesc: data.ProductDescription,
		Images:      data.ProductImages,
		Price:       data.ProductPrice,
	}

	// Store product in the database
	if err := models.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Scan the created product ID from the database
	var productID int
	if err := db.QueryRow("SELECT product_id FROM products ORDER BY product_id DESC LIMIT 1").Scan(&productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("productID", productID)
	producer.SendMessage(productID)

	// Return response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

}
