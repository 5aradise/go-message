package handlers

import (
	"net/http"

	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/middleware"
	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
)

func Signout(uDB types.UserGetterByName, uDelete types.UserDeleter) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := middleware.GetUser(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		auth.UnsetAuthCookie(c)
		auth.UnsetRefreshCookie(c)

		if user.ChatName != "" {
			err = uDelete.DeleteFromChat(user.ChatName, user.Name)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		}

		c.String(http.StatusOK, "")
	}
}
