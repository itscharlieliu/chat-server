package main

import (
	"log"
	"net/http"

	"github.com/itscharlieliu/chat-server/pkg"
)

func main() {

	c := make(chan []byte)

	http.HandleFunc("/chat", func(rw http.ResponseWriter, r *http.Request) { pkg.WebsocketHandler(rw, r, c) })

	go pkg.ChatHandler(c)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
