package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NicoMartina/nico-uretek-product-service/models"
	"github.com/go-chi/chi/v5"
)

var products []models.Product
var nextId int = 1

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
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

func GetProductByID(w http.ResponseWriter, r *http.Request) {
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

	for i := range products {
		if products[i].Id == id {
			if update.Name != "" {
				products[i].Name = update.Name
			}
			if update.Price > 0 {
				products[i].Price = update.Price
			}
			products[i].Description = update.Description

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}
