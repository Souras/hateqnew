package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
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
