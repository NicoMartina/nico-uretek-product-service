package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

type Product struct {
	Id int   `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}

var products []Product
var nextId int = 1

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    var p Product
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil || p.Name == "" || p.Price <= 0 {
        http.Error(w, "invalid input", http.StatusBadRequest)
        return
    }

    p.Id = nextId
    nextId++
    products = append(products, p)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err!= nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	var updated Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.Id == id {
			products[i].Name = updated.Name
			products[i].Price = updated.Price
			products[i].Description = updated.Description
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	r.Get("/products", getProducts)
	r.Post("/products", createProduct)
	r.Get("/products/{id}", getProductByID)
	r.Put("/products/{id}", updateProduct)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}