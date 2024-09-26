package ws

import (
	"log"
	"net/http"
	"slices"

	"github.com/5aradise/go-message/internal/middleware"
	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleNewConn(c *gin.Context) {
	user, err := middleware.GetUser(c)
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

	ChatUsers[user.Name] = conn

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			delete(ChatUsers, user.Name)
			break
		}

		if slices.Contains(message, senderMsgDiv) {
			continue
		}

		broadcastCh <- types.Message{Sender: user.Name, Body: message}
	}
}
