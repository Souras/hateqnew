func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Handle upgrade error
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			// Handle read error or connection close
			break
		}

		// Process the received message
		// ... (e.g., echo back the message)

		// Send a response message
		err = conn.WriteMessage(msgType, []byte("Received! "+string(message)))
		if err != nil {
			// Handle write error
			break
		}
	}
}