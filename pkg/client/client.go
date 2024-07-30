package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
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
	interrupt := make(chan os.Signal, 1)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("yo bitch whats your name ? ")
	clientName, _ := reader.ReadString('\n')
	//TODO: send name to server marshalling
	fmt.Println("Fuck you ", clientName)

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

	for {

		msg := []byte("ping")
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		default:
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("error pinging: ", err)
				return
			}
		}
		time.Sleep(time.Second * 1)

	}
}

func main() {
	Client()
}
