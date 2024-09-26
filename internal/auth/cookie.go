package auth

import (
	"github.com/gin-gonic/gin"
)

const (
	AccessTokenPath = "ACCESS_TOKEN"
	RefreshTokenPath = "REFRESH_TOKEN"
	UserNamePath = "name"
)

var (
	authMaxAge    = 15 * 60
	refreshMaxAge = 30 * 24 * 60 * 60
)

func SetAuthAndRefreshMaxAgeInSec(auth, refresh int) {
	authMaxAge = auth
	refreshMaxAge = refresh
}

func SetAuthCookie(c *gin.Context, accessToken, userName string) {
	c.SetCookie(AccessTokenPath, accessToken, authMaxAge, "/", "", false, true)
	c.SetCookie(UserNamePath, userName, authMaxAge, "/", "", false, false)
}

func SetRefreshCookie(c *gin.Context, refreshToken string) {
	c.SetCookie(RefreshTokenPath, refreshToken, refreshMaxAge, "/", "", false, true)
}

func UnsetAuthCookie(c *gin.Context) {
	c.SetCookie(AccessTokenPath, "", -1, "/", "", false, true)
	c.SetCookie(UserNamePath, "", -1, "/", "", false, false)
}

func UnsetRefreshCookie(c *gin.Context) {
	c.SetCookie(RefreshTokenPath, "", -1, "/", "", false, true)
}

func GetAccessToken(c *gin.Context) (string, error) {
	return c.Cookie(AccessTokenPath)
}

func GetRefreshToken(c *gin.Context) (string, error) {
	return c.Cookie(RefreshTokenPath)
}
