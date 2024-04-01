package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Souras/hateqnew/service/hateqService/internal/models"
)

func TestProducts(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{ID: 1, Name: "Product 1", Price: 10.50},
		{ID: 2, Name: "Product 2", Price: 20.75},
		{ID: 3, Name: "Product 2", Price: 22.75},
		{ID: 4, Name: "Product 2", Price: 22.75},
	}
	json.NewEncoder(w).Encode(products)
	// fmt.Fprintf(w, `{"message": "This is the API endpoint 2"}`)
}

// func ApiHandler(w http.ResponseWriter, r *http.Request) {
// 	// Handle API requests here
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprintf(w, `{"message": "This is the API endpoint 2"}`)
// }
