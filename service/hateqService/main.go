package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Souras/hateqnew/service/hateqService/internal/db"
	"github.com/Souras/hateqnew/service/hateqService/internal/routes"
)

func main() {

	err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// // Create a new router
	// r := mux.NewRouter()

	// // Register WebSocket handler
	// r.HandleFunc("/ws", handlers_common.websocketHandler)

	// // Register API handler
	// r.HandleFunc("/api", handlers_common.apiHandler)

	// r.HandleFunc("/", handlers_doctor.GetProducts).Methods("GET")
	// r.HandleFunc("/products/{id}", handlers_doctor.GetProduct).Methods("GET")
	// r.HandleFunc("/products", handlers_doctor.CreateProduct).Methods("POST")
	// r.HandleFunc("/products/{id}", handlers_doctor.UpdateProduct).Methods("PUT")

	// // Serve static files
	// fs := http.FileServer(http.Dir("static"))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r := routes.SetupRoutes()
	// Start HTTP server
	fmt.Println("Server listening on :5000")
	http.ListenAndServe(":5000", r)
}
