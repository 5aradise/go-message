package handlers

import (
	"net/http"
	"unicode/utf8"

	"github.com/5aradise/go-message/internal/types"
	"github.com/5aradise/go-message/pkg/random"
	"github.com/gin-gonic/gin"
)

type registerReq struct {
	Name string `json:"name" binding:"required"`
}

type registerResp struct {
	Key string `json:"key"`
}

func Register(uDB types.UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req registerReq
		if err := c.BindJSON(&req); err != nil {
			return
		}

		if !utf8.ValidString(req.Name) {
			c.String(http.StatusBadRequest, "Error: invalid utf8 string")
			return
		}

		_, ok := uDB.GetUserByName(req.Name)
		if ok {
			c.String(http.StatusBadRequest, "Error: user with this name already exist")
			return
		}

		key, err := random.String(64)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err.Error())
			return
		}

		uDB.SetUser(req.Name, &types.User{
			Name: req.Name,
			Key:  key,
		})

		c.JSON(http.StatusCreated, &registerResp{
			Key: key,
		})
	}
}
