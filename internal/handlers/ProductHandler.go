package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct{}

func (*ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8081/warehouses/1/stock/%s", mux.Vars(r)["id"]))
	if err != nil {
		http.Error(w, "Failed to make HTTP request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
	w.Write([]byte(fmt.Sprintf("This is the API product. You want: %s", mux.Vars(r)["id"])))
}
