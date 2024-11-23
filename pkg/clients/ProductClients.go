package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/CristianQS/otel-workshop-go/pkg/common"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ProductHandler struct{}

func (*ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	tracerProvider := common.ConsoleExporter(time.Second)
	tracer := tracerProvider.Tracer("store-api", trace.WithInstrumentationVersion("0.0.1"))
	_, span := tracer.Start(r.Context(), "get-product")
	defer span.End()
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
	productValueAttr := attribute.String("product.value", mux.Vars(r)["id"])
	warehouseValueAttr := attribute.Int("warehouse.value", 1)
	span.SetAttributes(productValueAttr)
	span.SetAttributes(warehouseValueAttr)

}
