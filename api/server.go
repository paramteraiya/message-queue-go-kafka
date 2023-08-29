package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/product", CreateProductHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}
