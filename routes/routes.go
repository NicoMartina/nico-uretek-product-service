package routes

import (
	"encoding/json"
	"net/http"

	"github.com/NicoMartina/nico-uretek-product-service/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Get("/products", handlers.GetProducts)
	r.Post("/products", handlers.CreateProduct)
	r.Get("/products/{id}", handlers.GetProductByID)
	r.Put("/products/{id}", handlers.UpdateProduct)

	return r
}