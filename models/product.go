package models

import (
	"github.com/lib/pq"
	"github.com/message-queue/database"
)


type Product struct {
	ID 					int
	// UserID 				int
	ProductName 		string
	ProductDesc 		string
	Images 				[]string
	Price 				float64
	CompressedImages	[]string
}

func CreateProduct (product *Product) error {
	db := database.GetDB()

	imagesArray := pq.Array(product.Images)

	query := "INSERT INTO products (product_name, product_desciption, product_images, product_price) VALUES ($1, $2, $3, $4)"
    _ = db.QueryRow(query, product.ProductName, product.ProductDesc, imagesArray, product.Price)


	// var productID int
	// err := row.Scan(&productID)
	// if err != nil {
	// 	return err
	// }

	// product.ID = productID
	return nil
}

func GetImages(productID int) []string {
	db := database.GetDB()
	var images []string
	rows, err := db.Query("SELECT product_images FROM products WHERE product_id = $1", productID)
	if err != nil {
		return images
	}
	for rows.Next() {
		var image string
		if err := rows.Scan(&image); err != nil {
			return nil
		}
		images = append(images, image)
	}
	return images
}

func StoreCompressedImages(productID int, compressedImages []string) error {
	db := database.GetDB()
	imagesArray := pq.Array(compressedImages)
	query := "UPDATE products SET compressed_product_images = $1, updated_at = now() WHERE product_id = $2"
	_, err := db.Exec(query, imagesArray, productID)
	return err
}