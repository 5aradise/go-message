package handlers

import (
	"log"
	"net/http"

	"github.com/5aradise/go-message/internal/types"
	"github.com/5aradise/go-message/internal/ws"
	"github.com/gin-gonic/gin"
)

func Signout(uDB types.UserGetterByWsKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		wsKey, err := c.Cookie("ws_key")
		if err != nil {
			c.String(http.StatusUnauthorized, "ws_key not found in cookies")
			return
		}

		u, err := uDB.GetUserByWebsocketKey(wsKey)
		if err != nil {
			c.String(http.StatusUnauthorized, "wrong ws_key")
			return
		}

		uConn, ok := ws.ChatUsers[u.Name]
		if ok {
			err := uConn.Close()
			if err != nil {
				log.Println("signout: comm.close:", err)
				c.String(http.StatusBadRequest, "user connection not found in chat")
				return
			}
			delete(ws.ChatUsers, u.Name)
		}

		c.SetCookie("ws_key", "", -1, "/", "", false, true)
		c.SetCookie("name", "", -1, "/", "", false, false)
		c.String(http.StatusOK, "")
	}
}
