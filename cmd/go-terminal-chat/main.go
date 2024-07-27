package main

import (
	"github.com/Sabir222/go-terminal-chat/pkg/client"
	"github.com/Sabir222/go-terminal-chat/pkg/server"
	"log"
	"net/http"
)

func main() {

	go func() {
		http.HandleFunc("/", server.Handler)
		log.Println("Server started on localhost:8080")
		// ":8080" instead of *addr
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	go func() {
		client.Client()
	}()

	select {}
}
