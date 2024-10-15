package ws

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type signal struct{}

type Server struct {
	chats map[string]*chat
	mu    sync.RWMutex
}

func NewServer() *Server {
	s := &Server{
		chats: make(map[string]*chat),
		mu:    sync.RWMutex{},
	}
	return s
}

func (s *Server) CreateChat(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.chats[name]
	if ok {
		return errors.New("chat with this name already on server")
	}

	c := NewChat()
	s.chats[name] = c
	go c.Run()
	return nil
}

func (s *Server) ConnectToChat(chatName, userName string, w http.ResponseWriter, r *http.Request) error {
	chat, ok := s.CheckChat(chatName)
	if !ok {
		return fmt.Errorf("chat with name %s unfound", chatName)
	}
	return chat.CreateUser(userName, w, r)
}

func (s *Server) DeleteFromChat(chatName, userName string) error {
	chat, ok := s.CheckChat(chatName)
	if !ok {
		return fmt.Errorf("chat with name %s unfound", chatName)
	}

	chat.DeleteUser(userName)
	return nil
}

func (s *Server) CheckChat(name string) (*chat, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ch, ok := s.chats[name]
	return ch, ok
}

func (s *Server) DeleteChat(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.chats[name]
	if !ok {
		log.Println("unfound chat by name:", name)
		return
	}

	c.stopCh <- signal{}
	for name := range c.users {
		c.DeleteUser(name)
	}
	close(c.broadCh)
	close(c.stopCh)

	delete(s.chats, name)
}
