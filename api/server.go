package api

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/product", CreateProductHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
