package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request, hub *ClientHub) {
	log.Println("Connected")

	upgrader.CheckOrigin = func(r *http.Request) bool {
		// TODO Sanitize origin
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := Client{
		hub:  hub,
		conn: conn,
		test: "hello",
	}

	client.hub.Register <- &client

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// channel <- p
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
