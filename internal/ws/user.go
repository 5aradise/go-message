package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/5aradise/go-message/internal/types"
	"github.com/gorilla/websocket"
)

type user struct {
	chat       *chat
	name       string
	conns      []*websocket.Conn
	connsCount uint8
	ch         chan types.Message
}

func NewUser(name string, ch *chat) *user {
	u := &user{
		name:  name,
		conns: make([]*websocket.Conn, 2),
		ch:    make(chan types.Message, 5),
		chat:  ch,
	}
	return u
}

func (u *user) AddConn(w http.ResponseWriter, r *http.Request, upgrader websocket.Upgrader) (id int, err error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return 0, err
	}

	return u.appendConn(conn), nil
}

func (u *user) ListenConn(id int) {
	conn := u.conns[id]
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			u.deleteConn(id)
			if u.connsCount == 0 {
				u.chat.DeleteUser(u.name)
			}
			return
		}
		u.chat.broadCh <- types.NewMessage(u.name, string(msg), time.Now())
	}
}

func (u *user) ListenChat() {
	for msg := range u.ch {
		err := u.writeMessage(msg)
		if err != nil {
			u.chat.DeleteUser(u.name)
			return
		}
	}
}

func (u *user) WriteSliceMsgsToConn(connId int, msgs []types.Message) {
	conn := u.conns[connId]
	if conn == nil {
		log.Println("WriteSliceMsgsToConn: try to write to nil conn")
		return
	}

	err := conn.WriteJSON(msgs)
	if err != nil {
		u.deleteConn(connId)
	}
	if u.connsCount == 0 {
		u.chat.DeleteUser(u.name)
	}
}

func (u *user) appendConn(newConn *websocket.Conn) (id int) {
	if len(u.conns) != int(u.connsCount) {
		for i, conn := range u.conns {
			if conn == nil {
				u.conns[i] = newConn
				u.connsCount++
				return i
			}
		}
	}
	u.conns = append(u.conns, newConn)
	u.connsCount++
	return int(u.connsCount - 2)
}

func (u *user) deleteConn(i int) {
	u.conns[i].Close()
	u.conns[i] = nil
	u.connsCount--
}

func (u *user) writeMessage(msg types.Message) error {
	var errToReturn error = nil
	for i, conn := range u.conns {
		if conn != nil {
			err := conn.WriteJSON(msg)
			if err != nil {
				u.deleteConn(i)
				errToReturn = err
			}
		}
	}
	if u.connsCount == 0 {
		return errToReturn
	}
	return nil
}
