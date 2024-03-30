package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sourabh/db"
	"github.com/sourabh/handlers"
)

func main() {
	err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	// r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
