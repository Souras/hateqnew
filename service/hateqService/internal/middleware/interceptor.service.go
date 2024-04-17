package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin //http://localhost:8100
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow the POST method
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's a preflight request, respond with success status
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func AddResponseHeader(w http.ResponseWriter) http.ResponseWriter {
	// Set common headers here
	w.Header().Set("Content-Type", "application/json") // Example header
	w.Header().Set("Access-Control-Allow-Origin", "*") // Example header (adjust as needed)
	return w
}

func AddResponseHeaderWithoutAuth(h http.Handler) http.Handler {
	// Create a handler that wraps the original handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set common headers here
		w.Header().Set("Content-Type", "application/json") // Example header
		w.Header().Set("Access-Control-Allow-Origin", "*") // Example header (adjust as needed)

		// Call the original handler
		h.ServeHTTP(w, r)
	})
}

func yourHandlerFunction(w http.ResponseWriter, r *http.Request) {
	// ... (your request handling logic)
	// Headers are already set by the middleware
}
