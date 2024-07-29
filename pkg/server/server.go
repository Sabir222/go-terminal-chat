package main

import (
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Printf("error writing message: %s  ", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Server started on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
