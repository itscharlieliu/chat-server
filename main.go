package main

import (
	"log"
	"net/http"

	"github.com/itscharlieliu/chat-server/pkg"
)

func main() {

	hub := pkg.ClientHub{
		ClientMap:  make(map[*pkg.Client]bool),
		Send:       make(chan []byte),
		Register:   make(chan *pkg.Client),
		Deregister: make(chan *pkg.Client),
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { pkg.WebsocketHandler(rw, r, &hub) })

	go pkg.ChatHandler(&hub)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
