# go-message

## Description

A simple web messenger written in Go programming language. 

It uses [Gin](https://gin-gonic.com/) as the HTTP framework with [gorilla/websocket](https://github.com/gorilla/websocket) for real-time communication and [SQLite](https://www.sqlite.org/) as the database with [gorm](https://gorm.io/) as the ORM library.

## Features
- Real-time messaging
- Authentication and authorization

## Technologies
- Go
- Gin
- Gorilla WebSocket
- Gorm
- SQLite

## Requirements
- Go 1.23
- SQLite

## Running the application
1. Clone the repository
2. Create a copy of the `.env.example` file and rename it to `.env`
3. Run `make run` to start the server
