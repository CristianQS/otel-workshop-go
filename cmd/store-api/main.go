package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	productHandler := ProductHandler{}
	r.HandleFunc("/products/{id}", productHandler.ServeHttp)
	http.ListenAndServe(":8080", r)
}

type ProductHandler struct{}

func (*ProductHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is the API product.You want: %s", mux.Vars(r)["id"])))
}
