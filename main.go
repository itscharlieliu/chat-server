package main

import (
	"log"
	"net/http"

	"github.com/itscharlieliu/chat-server/pkg"
)

func main() {

	http.HandleFunc("/chat", pkg.WebsocketHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
