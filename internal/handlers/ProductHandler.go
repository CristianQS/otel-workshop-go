package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct{}

func (*ProductHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is the API product.You want: %s", mux.Vars(r)["id"])))
}
