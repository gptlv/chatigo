package ws

import "fmt"

type Server struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	rooms      map[*Room]bool
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		rooms:      make(map[*Room]bool),
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.registerClient(client)
		case client := <-s.unregister:
			s.unregisterClient(client)
		case message := <-s.broadcast:
			s.broadcastToClients(message)
		}
	}
}

func (s *Server) registerClient(c *Client) {
	s.notifyClientJoined(c)
	s.listOnlineClients(c)
	s.clients[c] = true
}

func (s *Server) unregisterClient(c *Client) {
	if _, ok := s.clients[c]; ok {
		delete(s.clients, c)
		s.notifyClientLeft(c)
	}

}

func (s *Server) broadcastToClients(message []byte) {
	for c := range s.clients {
		c.send <- message
	}
}

func (s *Server) createRoom(name string, kind string) *Room {
	room := NewRoom(name, kind)
	go room.RunRoom()
	s.rooms[room] = true
	return room
}

const clientJoinedMessage = "%s just joined the chat!"

func (s *Server) notifyClientJoined(c *Client) {
	message := &Message{
		Message: fmt.Sprintf(clientJoinedMessage, c.GetName()),
		Action:  UserJoinedAction,
	}
	s.broadcastToClients(message.encode())

}

const clientLeftMessage = "%s left"

func (s *Server) notifyClientLeft(c *Client) {
	message := &Message{
		Message: fmt.Sprintf(clientJoinedMessage, c.GetName()),
		Action:  UserLeftAction,
	}
	s.broadcastToClients(message.encode())

}

func (s *Server) listOnlineClients(c *Client) {
	for existingClient := range s.clients {
		message := &Message{
			Action: UserJoinedAction,
			Sender: existingClient,
		}
		c.send <- message.encode()
	}
}

func (s *Server) findClientByID(ID string) *Client {
	var foundClient *Client

	for client := range s.clients {
		if client.GetID() == ID {
			foundClient = client
			break
		}
	}

	return foundClient
}

func (s *Server) findRoomByID(ID string) *Room {
	var foundRoom *Room
	for room := range s.rooms {
		if room.GetID() == ID {
			foundRoom = room
			break
		}
	}
	return foundRoom
}
