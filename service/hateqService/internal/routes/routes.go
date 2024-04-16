package routes

import (
	"net/http"

	handlers_doctor "github.com/Souras/hateqnew/service/hateqService/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// Register API handler
	router.HandleFunc("/login", middleware.LoginHandler).Methods("GET")
	// router.HandleFunc("/api", handlers_doctor.TestProducts).Methods("GET")
	router.Handle("/api", middleware.AuthMiddleware(http.HandlerFunc(handlers_doctor.TestProducts))).Methods("GET")
	router.HandleFunc("/ws", handlers_doctor.WebsocketHandler)

	router.HandleFunc("/", handlers_doctor.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers_doctor.GetProduct).Methods("GET")
	router.HandleFunc("/products", handlers_doctor.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers_doctor.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers_doctor.DeleteProduct).Methods("DELETE")

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

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}
