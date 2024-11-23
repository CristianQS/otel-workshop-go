package main

import (
	"net/http"

	"github.com/CristianQS/otel-workshop-go/pkg/clients"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	WarehouseHandler := clients.WarehouseHandler{}
	r.HandleFunc("/warehouses/{id}", WarehouseHandler.ServeHttp).Methods("GET")
	r.HandleFunc("/warehouses/{id}/stock/{productId}", WarehouseHandler.GetStock).Methods("GET")
	http.ListenAndServe(":8081", r)
}
