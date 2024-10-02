package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type WarehouseHandler struct{}

func (*WarehouseHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is the API warehouse. You want: %s", mux.Vars(r)["id"])))
}

func (*WarehouseHandler) GetStock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("The stock %s from warehouse %s", mux.Vars(r)["productId"], mux.Vars(r)["id"])))
}
