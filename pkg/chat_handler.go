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
	test string
}

type ClientHub struct {
	ClientMap  map[*Client]bool
	Send       chan []byte
	Register   chan *Client
	Deregister chan *Client
}

func ChatHandler(hub *ClientHub) {

	for {
		select {
		case client := <-hub.Register:
			log.Printf(client.test)
		}
	}
}

// TODO
// https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go
// https://stackoverflow.com/questions/32693931/what-is-the-benefit-of-sending-to-a-channel-by-using-select-in-go
// https://stackoverflow.com/questions/31532652/go-websocket-send-all-clients-a-message
