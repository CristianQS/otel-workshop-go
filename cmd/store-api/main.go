package main

import (
	"net/http"

	"github.com/CristianQS/otel-workshop-go/pkg/clients"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	productHandler := clients.ProductHandler{}
	r.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")

	http.ListenAndServe(":8080", r)
}
