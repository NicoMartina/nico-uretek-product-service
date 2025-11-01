package main

import (
	"log"
	"net/http"

	"github.com/NicoMartina/nico-uretek-product-service/routes"
)

func main() {
	r := routes.RegisterRoutes()

	

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}