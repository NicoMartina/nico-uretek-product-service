package store

import (
	"encoding/json"
	"log"
	"os"

	"github.com/NicoMartina/nico-uretek-product-service/models"
)

var Products []models.Product
var NextId int = 1

func LoadProducts() {
	file, err := os.Open("products.json")
	if err != nil {
		log.Println("No exisiting  product  file found, starting fresh.")
		return
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&Products)

	for _, p := range Products {
		if p.Id == NextId {
			NextId = p.Id + 1
		}
	}
}

func SaveProducts() {
	file, err := os.Create("products.json")
	if err != nil {
		log.Println("Error saving products:", err)
		return
	}

	defer file.Close()

	json.NewEncoder(file).Encode(Products)
}