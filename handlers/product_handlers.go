package handlers

import (
	"net/http"
	"strconv"

	"github.com/NicoMartina/nico-uretek-product-service/store"
	"github.com/go-chi/chi/v5"
)

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