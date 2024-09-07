package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/5aradise/go-message/internal/types"
	"github.com/5aradise/go-message/pkg/valid"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Register(uDB types.UserCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req registerReq
		if err := c.BindJSON(&req); err != nil {
			return
		}

		if err := valid.Name(req.Name); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if err := valid.Password(req.Password); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		email := sql.NullString{String: req.Email}
		if req.Email != "" {
			if err := valid.Email(req.Email); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			email.Valid = true
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, "internal error")
			return
		}

		_, err = uDB.CreateUser(req.Name, hashedPassword, email)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				if strings.Contains(err.Error(), "name") {
					c.String(http.StatusBadRequest, "user with this name already exists")
					return
				}
				c.String(http.StatusBadRequest, "user with this email already exists")
				return
			}
			log.Println(err)
			c.String(http.StatusBadRequest, "internal error")
			return
		}

		c.String(http.StatusCreated, "user successfuly created")
	}
}
