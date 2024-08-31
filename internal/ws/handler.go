package ws

import (
	"log"
	"slices"

	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleNewConn(uDB types.UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("newConn: upgrade:", err)
			return
		}
		defer conn.Close()

		_, key, err := conn.ReadMessage()
		if err != nil {
			log.Println("newConn: read:", err)
			return
		}

		u, ok := uDB.GetUserByKey(string(key))
		if !ok {
			log.Println("newConn: error: get user by key")
			return
		}

		u.WsConn = conn

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				uDB.DeleteUser(u.Name)
				log.Println("newConn: read:", err)
				break
			}

			if slices.Contains(message, senderMsgDiv) {
				continue
			}

			broadcastCh <- types.Message{Sender: u.Name, Body: message}
		}
	}
}
