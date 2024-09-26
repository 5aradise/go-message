package handlers

import (
	"log"
	"net/http"

	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/middleware"
	"github.com/5aradise/go-message/internal/types"
	"github.com/5aradise/go-message/internal/ws"
	"github.com/gin-gonic/gin"
)

func Signout(uDB types.UserGetterByName) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := middleware.GetUser(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		uConn, ok := ws.ChatUsers[user.Name]
		if ok {
			err := uConn.Close()
			if err != nil {
				log.Println("signout: comm.close:", err)
				c.String(http.StatusBadRequest, "user connection not found in chat")
				return
			}
			delete(ws.ChatUsers, user.Name)
		}

		auth.UnsetAuthCookie(c)
		auth.UnsetRefreshCookie(c)
		c.String(http.StatusOK, "")
	}
}
