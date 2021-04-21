package pkg

import (
	"github.com/gorilla/websocket"
)

type ThreadsafeConn struct {
	conn websocket.Conn
	msg  chan []byte
}

type ConnHub struct {
	connMap map[*ThreadsafeConn]bool
}

func ChatHandler(msgChannel chan []byte, connChannel chan ThreadsafeConn) {

	hub := connHub{connMap: make(map[*ThreadsafeConn]bool)}

	for {
		msg := <-msgChannel

		for threadsafeConn := range hub.connMap {
			threadsafeConn.conn.WriteMessage(1, msg)
		}
	}
}

// TODO
// https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go
// https://stackoverflow.com/questions/32693931/what-is-the-benefit-of-sending-to-a-channel-by-using-select-in-go
// https://stackoverflow.com/questions/31532652/go-websocket-send-all-clients-a-message
