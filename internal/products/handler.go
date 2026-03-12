package products

import (
	"log"
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
	// ~ so service is going to defined over there
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products := []string{"Hello", "World"}
	json.WriteJson(w, http.StatusOK, products)
}
