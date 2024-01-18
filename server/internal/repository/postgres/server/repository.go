package server

import (
	"github.com/gptlv/chatigo/server/internal/domain"
)

type serverRepository struct {
	users      map[*domain.User]bool
	register   chan *domain.User
	unregister chan *domain.User
	broadcast  chan []byte
	rooms      map[*domain.Room]bool
}

func NewServerRepository() domain.ServerRepository {
	return &serverRepository{
		users:      make(map[*domain.User]bool),
		register:   make(chan *domain.User),
		unregister: make(chan *domain.User),
		broadcast:  make(chan []byte),
		rooms:      make(map[*domain.Room]bool),
	}
}
