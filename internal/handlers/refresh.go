package handlers

import (
	"net/http"

	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
)

func Refresh(uDB types.UserGetterByRefreshToken, jwtC types.JWTCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("REFRESH_TOKEN")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "REFRESH_TOKEN not found in cookies"})
			return
		}

		u, err := uDB.GetUserByRefreshToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong refresh token"})
			return
		}

		accToken, err := jwtC.CreateJWTtoken(u.Name)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		auth.SetAuthCookie(c, accToken, u.Name)
		c.String(http.StatusOK, "")
	}
}
