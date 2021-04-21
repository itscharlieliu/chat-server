package main

import (
	"log"
	"net/http"

	"github.com/itscharlieliu/chat-server/pkg"
)

func main() {

	msgChannel := make(chan []byte)
	connChannel := make(chan pkg.ThreadsafeConn)

	http.HandleFunc("/chat", func(rw http.ResponseWriter, r *http.Request) { pkg.WebsocketHandler(rw, r, msgChannel) })

	go pkg.ChatHandler(msgChannel)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
