package main

import (
	"net/http"

	"github.com/CristianQS/otel-workshop-go/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	productHandler := handlers.ProductHandler{}
	r.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")

	http.ListenAndServe(":8080", r)
}
