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
	}

	client.hub.Register <- &client

	defer func() {
		client.hub.Deregister <- &client
		client.conn.Close()
	}()

	for {
		messageType, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		msg := Message{
			bytes:       bytes,
			messageType: messageType,
		}

		hub.Send <- msg
		// // channel <- p
		// err = conn.WriteMessage(messageType, p)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}
