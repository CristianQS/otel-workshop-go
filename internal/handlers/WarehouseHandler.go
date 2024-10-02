package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type WarehouseHandler struct{}

func (*WarehouseHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is the API product.You want: %s", mux.Vars(r)["id"])))
}
