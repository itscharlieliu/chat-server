package main

import (
	"log"
	"net/http"

	"github.com/itscharlieliu/chat-server/pkg"
)

func main() {

	channel := make(chan []byte)

	http.HandleFunc("/chat", func(rw http.ResponseWriter, r *http.Request) { pkg.WebsocketHandler(rw, r, channel) })

	go pkg.ChatHandler(channel)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
