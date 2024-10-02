package main

import (
	"net/http"

	"github.com/CristianQS/otel-workshop-go/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	WarehouseHandler := handlers.WarehouseHandler{}
	r.HandleFunc("/warehouses/{id}", WarehouseHandler.ServeHttp)
	http.ListenAndServe(":8081", r)
}
