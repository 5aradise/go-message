package ws

import (
	"log"
	"net/http"
	"slices"

	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleNewConn(uDB types.UserGetterByWsKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		wsKey, err := c.Cookie("ws_key")
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		u, err := uDB.GetUserByWebsocketKey(wsKey)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("newConn: upgrade:", err)
			return
		}
		defer conn.Close()

		ChatUsers[u.Name] = conn

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				delete(ChatUsers, u.Name)
				break
			}

			if slices.Contains(message, senderMsgDiv) {
				continue
			}

			broadcastCh <- types.Message{Sender: u.Name, Body: message}
		}
	}
}
