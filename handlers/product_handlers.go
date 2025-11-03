package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NicoMartina/nico-uretek-product-service/models"
	"github.com/NicoMartina/nico-uretek-product-service/store"
	"github.com/go-chi/chi/v5"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.Products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil || p.Name == "" || p.Price <= 0 {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	p.Id =store.NextId
	store.NextId++
	store.Products = append(store.Products, p)
	store.SaveProducts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid  product ID", http.StatusBadRequest)
		return
	}
	for _, p := range store.Products {
		if p.Id == id {
			w.Header().Set("Content-Type", "application.json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	var update models.Product
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	for i := range store.Products {
		if store.Products[i].Id == id {
			if update.Name != "" {
				store.Products[i].Name = update.Name
			}
			if update.Price > 0 {
				store.Products[i].Price = update.Price
			}
			store.Products[i].Description = update.Description

			store.SaveProducts()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(store.Products[i])
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	for i, p := range store.Products {
		if p.Id == id {
			store.Products = append(store.Products[:i], store.Products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}