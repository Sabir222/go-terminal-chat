package client

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func Client() {

	url := "ws://localhost:8080"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial failed")
	}

	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("error reading incomming messges", err)
			}
			log.Println("Received: in the client loop", string(message))
		}
	}()

	for {
		msg := []byte("Hello from client")
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error sending msg")
		}
		time.Sleep(1 * time.Second)
	}

}
