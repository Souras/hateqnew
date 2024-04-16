package middleware

import (
	"net/http"
)

func CorsMiddleware(h http.Handler) http.Handler {
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

// func main() {
//   // Wrap your main handler with the middleware
//   handler := corsMiddleware(http.HandlerFunc(yourHandlerFunction))

//   // Start your server
//   http.ListenAndServe(":8080", handler)
// }
