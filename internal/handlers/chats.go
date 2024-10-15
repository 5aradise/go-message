package handlers

import (
	"net/http"

	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/middleware"
	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
)

type createChatReq struct {
	Name string `json:"name" binding:"required"`
}

func CreateChat(chCr types.ChatCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := middleware.GetUser(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		var req createChatReq
		if err := c.BindJSON(&req); err != nil {
			return
		}

		err = chCr.CreateChat(req.Name)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusCreated, "chat with name %s successfully created", req.Name)
	}
}

func ConnectToChat(chConn types.ChatConnecter) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := middleware.GetUser(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		chatName := c.Param("chatName")
		if chatName == "" {
			c.String(http.StatusBadRequest, "empty chat name")
			return
		}

		err = chConn.ConnectToChat(chatName, user.Name, c.Writer, c.Request)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		auth.SetChatCookie(c, chatName)

		c.String(http.StatusOK, "")
	}
}
