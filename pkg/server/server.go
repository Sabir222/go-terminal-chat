package server

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader is a websocket.Upgrader that is used to upgrade an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection.
	// `w` is the HTTP response writer, and `r` is the HTTP request.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Log an error if the upgrade fails.
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	go readLoop(conn)

	for {
		messageType, messageReader, err := conn.NextReader()
		if err != nil {
			log.Printf("Error getting message reader: %v", err)
			break
		}

		messageWriter, err := conn.NextWriter(messageType)
		if err != nil {
			log.Printf("Error getting message writer: %v", err)
			break
		}

		// Copy the data from the message reader to the message writer.
		// This effectively echoes the received message back to the client.
		if _, err := io.Copy(messageWriter, messageReader); err != nil {
			// Log an error if copying the message fails.
			log.Printf("Error copying message: %v", err)
			break
		}

		// Close the message writer to finalize the WebSocket frame.
		if err := messageWriter.Close(); err != nil {
			// Log an error if closing the message writer fails.
			log.Printf("Error closing message writer: %v", err)
			break
		}
	}
}

func readLoop(c *websocket.Conn) {
	for {
		_, msg, err := c.NextReader()
		if err != nil {
			c.Close()
			break
		}

		message, err := io.ReadAll(msg)
		if err != nil {
			log.Fatal("shit")
		}
		log.Printf("Message content: %s", message)
	}
}
