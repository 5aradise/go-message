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

	db := database.New()

	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Secure(""),
	)

	// main
	r.StaticFile("/", "./public/index.html")

	// static
	r.StaticFile("/favicon.ico", "./public/images/favicon.ico")
	r.Static("/static", "./public")

	// api
	api := r.Group("/api")

	api.GET("/ping", handlers.Ping)
	api.POST("/register", handlers.Register(db))

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
	go ws.RunBroadcast(db)
	err = srv.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
