package handlers

import (
	"net/http"

	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/types"
	"github.com/5aradise/go-message/pkg/random"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(uDB interface {
	types.UserGetterByName
	types.RefreshTokenUpdater
}, jwtC types.JWTCreator) gin.HandlerFunc {
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

		refreshToken, err := random.String(64)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		err = uDB.UpdateRefreshTokenByUserName(u.Name, refreshToken)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		accToken, err := jwtC.CreateJWTtoken(u.Name)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		auth.SetAuthCookie(c, accToken, u.Name)
		auth.SetRefreshCookie(c, refreshToken)
		c.String(http.StatusOK, "")
	}
}
