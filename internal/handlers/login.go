package handlers

import (
	"net/http"

	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const deyInSec = 24*60*60

type loginReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(uDB types.UserGetterByName) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginReq
		if err := c.BindJSON(&req); err != nil {
			return
		}

		u, err := uDB.GetUserByName(req.Name)
		if err != nil {
			c.String(http.StatusBadRequest, "wrong password")
			return
		}

		err = bcrypt.CompareHashAndPassword(u.Password, []byte(req.Password))
		if err != nil {
			c.String(http.StatusBadRequest, "wrong password")
			return
		}

		c.SetCookie("ws_key", u.WebsocketKey, deyInSec, "/", "", false, true)
		c.SetCookie("name", u.Name, deyInSec, "/", "", false, false)
		c.String(http.StatusOK, "")
	}
}
