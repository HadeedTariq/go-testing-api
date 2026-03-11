package products

import (
	"net/http"

	"github.com/HadeedTariq/go-testing-api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := []string{"Hello", "World"}
	json.WriteJson(w, http.StatusOK, products)
}
