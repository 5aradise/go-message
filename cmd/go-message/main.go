package main

import (
	"log"
	"net"
	"net/http"
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

	jwtService := jwt.New(cfg.JWT.Key, APP_NAME, time.Duration(cfg.Auth.AccessTokenMaxAge+5)*time.Second)

	r := gin.New()
	auth.SetAuthAndRefreshMaxAgeInSec(cfg.Auth.AccessTokenMaxAge, cfg.Auth.RefreshTokenMaxAge)

	authMid := middleware.Auth(jwtService, db)

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Secure(""),
	)

	// static
	r.StaticFile("/", "./public/index.html")
	r.StaticFile("/signup", "./public/signup.html")
	r.StaticFile("/login", "./public/login.html")

	r.StaticFile("/favicon.ico", "./public/images/favicon.ico")

	r.Static("/static", "./public")

	// api
	api := r.Group("/api")

	api.GET("/ping", handlers.Ping)
	api.POST("/register", handlers.Register(db))
	api.POST("/login", handlers.Login(db, jwtService))

	api.POST("/signout", authMid, handlers.Signout(db))

	api.POST("/refresh", handlers.Refresh(db, jwtService))

	api.GET("/ws", authMid, ws.HandleNewConn)

	// no route
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/404.html")
	})

	srv := &http.Server{
		Addr:              net.JoinHostPort("", cfg.Server.Port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Starting HTTP server on port %s", cfg.Server.Port)
	go ws.RunBroadcast()
	err = srv.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
