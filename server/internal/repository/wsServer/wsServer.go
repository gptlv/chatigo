package wsserver

import "github.com/gptlv/chatigo/server/internal/domain"

type server struct {
	users      map[*User]bool
	register   chan *User
	unregister chan *User
	broadcast  chan []byte
	rooms      map[*Room]bool
}

func NewServer() *domain.WsServer {
	server := &server{
		users:      make(map[*domain.User]bool),
		register:   make(chan *domain.User),
		unregister: make(chan *domain.User),
		broadcast:  make(chan []byte),
		rooms:      make(map[*domain.Room]bool),
	}

	return server
}
