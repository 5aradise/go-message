package main

import (
	"log"
	"net"
	"net/http"
	"path/filepath"
	"time"

	"github.com/5aradise/go-message/config"
	"github.com/5aradise/go-message/internal/auth"
	"github.com/5aradise/go-message/internal/database"
	"github.com/5aradise/go-message/internal/handlers"
	"github.com/5aradise/go-message/internal/middleware"
	"github.com/5aradise/go-message/internal/ws"
	"github.com/5aradise/go-message/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const APP_NAME = "go-message"

func main() {
	godotenv.Load()

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg.DB.Path)
	if err != nil {
		log.Fatal(err)
	}

	auth.SetAuthAndRefreshMaxAgeInSec(cfg.Auth.AccessTokenMaxAge, cfg.Auth.RefreshTokenMaxAge)

	jwtService := jwt.New(cfg.JWT.Key, APP_NAME, time.Duration(cfg.Auth.AccessTokenMaxAge+5)*time.Second)

	authMid := middleware.Auth(jwtService, db)

	wsServer := ws.NewServer(cfg.WS.ReadBufferSize, cfg.WS.WriteBufferSize)

	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Secure(""),
	)

	publicPath := filepath.Join(".", "public")

	// static
	r.StaticFile("/", filepath.Join(publicPath, "index.html"))
	r.StaticFile("/signup", filepath.Join(publicPath, "signup.html"))
	r.StaticFile("/login", filepath.Join(publicPath, "login.html"))
	r.GET("/chats/:chatName",
		func(c *gin.Context) {
			chatName := c.Param("chatName")
			if chatName == "" {
				c.Redirect(http.StatusBadGateway, "/")
				return
			}
			c.File(filepath.Join(publicPath, "chat.html"))
		},
	)

	r.StaticFile("/favicon.ico", filepath.Join(publicPath, "images", "favicon.ico"))

	r.Static("/static", publicPath)

	// api
	api := r.Group("/api")

	api.GET("/ping", handlers.Ping)
	api.POST("/register", handlers.Register(db))
	api.POST("/login", handlers.Login(db, jwtService))

	api.POST("/signout", authMid, handlers.Signout(db, wsServer))

	api.POST("/refresh", handlers.Refresh(db, jwtService))

	api.POST("/chats", authMid, handlers.CreateChat(wsServer))
	// TODO api.DELETE("/chats/:chatName", handlers.DeleteChat())
	api.GET("/ws/:chatName", authMid, handlers.ConnectToChat(wsServer))

	// no route
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(publicPath, "404.html"))
	})

	srv := &http.Server{
		Addr:              net.JoinHostPort("", cfg.Server.Port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Starting HTTP server on port %s", cfg.Server.Port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
