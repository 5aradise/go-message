package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/5aradise/go-message/config"
	"github.com/5aradise/go-message/internal/database"
	"github.com/5aradise/go-message/internal/handlers"
	"github.com/5aradise/go-message/internal/middleware"
	"github.com/5aradise/go-message/internal/ws"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	r := gin.New()

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
	api.POST("/login", handlers.Login(db))
	api.POST("/signout", handlers.Signout(db))

	api.GET("/ws", ws.HandleNewConn(db))

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
