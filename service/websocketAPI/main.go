package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Echo back messages received from client
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		if string(p) != "" {
			fmt.Println("reading message:", string(p))
		}
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Handle API requests here
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "This is the API endpoint 2"}`)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register WebSocket handler
	r.HandleFunc("/ws", websocketHandler)

	// Register API handler
	r.HandleFunc("/api", apiHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Start HTTP server
	fmt.Println("Server listening on :5000")
	http.ListenAndServe(":5000", r)
}
