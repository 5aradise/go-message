package ws

import (
	"log"

	"github.com/5aradise/go-message/internal/types"
	"github.com/gorilla/websocket"
)

const senderMsgDiv = 0

var broadcastCh = make(chan types.Message, 10)
var ChatUsers = make(map[string]*websocket.Conn)

func RunBroadcast() {
	for msg := range broadcastCh {
		toSend := make([]byte, 0, len(msg.Sender)+len(msg.Body)+1)
		toSend = append(toSend, []byte(msg.Sender)...)
		toSend = append(toSend, senderMsgDiv)
		toSend = append(toSend, msg.Body...)
		for name, conn := range ChatUsers {
			if name != msg.Sender && conn != nil {
				err := conn.WriteMessage(websocket.BinaryMessage, toSend)
				if err != nil {
					log.Println("broadcast:", err)
				}
			}
		}
	}
}
