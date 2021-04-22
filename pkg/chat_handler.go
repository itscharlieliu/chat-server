package pkg

import (
	"log"

	"github.com/gorilla/websocket"
)

// type ThreadsafeConn struct {
// 	conn websocket.Conn
// }

type Client struct {
	hub  *ClientHub
	conn *websocket.Conn
}

type Message struct {
	bytes       []byte
	messageType int
}

type ClientHub struct {
	ClientMap  map[*Client]bool
	Send       chan Message
	Register   chan *Client
	Deregister chan *Client
}

func ChatHandler(hub *ClientHub) {

	for {
		select {
		case client := <-hub.Register:
			log.Println("Connected")
			hub.ClientMap[client] = true
		case client := <-hub.Deregister:
			log.Println("Disconnected")
			delete(hub.ClientMap, client)
		case msg := <-hub.Send:
			log.Println(string(msg.bytes))
			for client := range hub.ClientMap {
				err := client.conn.WriteMessage(msg.messageType, msg.bytes)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

// TODO
// https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go
// https://stackoverflow.com/questions/32693931/what-is-the-benefit-of-sending-to-a-channel-by-using-select-in-go
// https://stackoverflow.com/questions/31532652/go-websocket-send-all-clients-a-message
