package middleware

import (
	"errors"
	"net/http"

	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/types"
	"github.com/gin-gonic/gin"
)

func Auth(jwtS types.SubGetter, uDB types.UserGetterByName) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := auth.GetAccessToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "ACCESS_TOKEN not found in cookies"})
			return
		}

		sub, err := jwtS.GetSubjectFromJWT(jwt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid ACCESS_TOKEN: " + err.Error()})
			return
		}

		user, err := uDB.GetUserByName(sub)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user by ACCESS_TOKEN not found: " + err.Error()})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func GetUser(c *gin.Context) (types.User, error) {
	uu, ok := c.Get("user")
	if !ok {
		return types.User{}, errors.New("unauthorized user")
	}

	user, ok := uu.(types.User)
	if !ok {
		return types.User{}, errors.New("unauthorized user")
	}
	return user, nil
}
