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

	done := make(chan struct{})
	msgChannel := make(chan []byte)

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("error reading incomming messges", err)
			}
			log.Println("client", string(message))
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case msg, ok := <-msgChannel:
				if !ok {
					return
				}
				err := conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Printf("error sending msg")
				}
			}

		}
	}()

	go func() {
		for {
			msg := []byte("Ping!")
			msgChannel <- msg
			time.Sleep(time.Second * 1)
		}

	}()
	<-done

}
