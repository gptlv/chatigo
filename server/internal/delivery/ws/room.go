package ws

import (
	"fmt"

	"github.com/google/uuid"
)

const PrivateRoomKind = "private"
const PublicRoomKind = "public"

type Room struct {
	ID         uuid.UUID
	Kind       string
	name       string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

func NewRoom(name string, kind string) *Room {
	return &Room{
		ID:         uuid.New(),
		Kind:       kind,
		name:       name,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}

}

func (r *Room) RunRoom() {
	for {
		select {
		case client := <-r.register:
			r.registerClient(client)
		case client := <-r.unregister:
			r.unregisterClient(client)
		case message := <-r.broadcast:
			r.sendMessageToClientsInRoom(message.encode())
		}

	}
}

func (r *Room) registerClient(c *Client) {
	if r.Kind != PrivateRoomKind {
		r.sendWelcomeMessage(c)
	}
	r.clients[c] = true
}

func (r *Room) unregisterClient(c *Client) {
	if _, ok := r.clients[c]; ok {
		delete(r.clients, c)
	}
}

func (r *Room) sendMessageToClientsInRoom(m []byte) {
	for client := range r.clients {
		client.send <- m
	}
}

func (r *Room) GetName() string {
	return r.name
}

func (r *Room) SetName(name string) {
	r.name = name
}

const welcomeMessage = "%s just joined the room!"

func (r *Room) sendWelcomeMessage(c *Client) {
	message := &Message{
		Action:  SendMessageAction,
		Message: fmt.Sprintf(welcomeMessage, c.GetName()),
		Target:  r,
	}

	r.sendMessageToClientsInRoom(message.encode())
}

func (r *Room) GetID() string {
	return r.ID.String()
}
