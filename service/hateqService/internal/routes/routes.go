package routes

import (
	"net/http"

	"github.com/Souras/hateqnew/service/hateqService/internal/handlers/handlers_doctor"
	"github.com/gorilla/mux"
	// "your_project/handlers"
	// "your_project/middlewares"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// Middleware
	// router.Use(middlewares.LoggingMiddleware)
	// router.Use(middlewares.AuthenticationMiddleware)

	// Token-related routes
	// tokenRouter := router.PathPrefix("/tokens").Subrouter()
	// tokenRouter.HandleFunc("/", handlers.GenerateTokenHandler).Methods("POST")
	// tokenRouter.HandleFunc("/{tokenID}", handlers.GetTokenHandler).Methods("GET")

	// // Patient-related routes
	// patientRouter := router.PathPrefix("/patients").Subrouter()
	// patientRouter.HandleFunc("/", handlers.GetPatientsHandler).Methods("GET")

	// // Doctor-related routes
	// doctorRouter := router.PathPrefix("/doctors").Subrouter()
	// doctorRouter.HandleFunc("/{doctorID}/call", handlers.CallPatientHandler).Methods("POST")

	// Register WebSocket handler
	// router.HandleFunc("/ws", handlers_common.websocketHandler)

	// Register API handler
	router.HandleFunc("/api", handlers_common.apiHandler).Methods("GET")

	router.HandleFunc("/", handlers_doctor.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers_doctor.GetProduct).Methods("GET")
	router.HandleFunc("/products", handlers_doctor.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers_doctor.UpdateProduct).Methods("PUT")

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}
