package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

type Client struct {
	ID     uuid.UUID
	name   string
	conn   *websocket.Conn
	server *Server
	send   chan []byte
	rooms  map[*Room]bool
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func newClient(conn *websocket.Conn, server *Server, name string) *Client {
	return &Client{
		ID:     uuid.New(),
		name:   name,
		conn:   conn,
		server: server,
		send:   make(chan []byte),
		rooms:  make(map[*Room]bool),
	}
}

func (c *Client) readPump() {
	defer func() {
		c.disconnect()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, jsonMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		c.handleNewMessage(jsonMessage)
	}

}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("writer error: %v", err)
				return
			}

			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Printf("close writer error: %v", err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("write message error: %v", err)
				return
			}
		}

	}
}

func (c *Client) disconnect() {
	c.server.unregister <- c
	close(c.send)
	c.conn.Close()
	for room := range c.rooms {
		room.unregister <- c
	}
}

func (c *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message: %v", err)
	}

	message.Sender = c

	switch message.Action {
	case SendMessageAction:
		room := message.Target

		if room := c.server.findRoomByID(room.GetID()); room != nil {
			room.broadcast <- &message
		}

	case JoinRoomAction:
		c.handleJoinRoomMessage(message)

	case LeaveRoomAction:
		c.handleLeaveRoomMessage(message)

	case JoinPrivateRoomAction:
		c.handleJoinPrivateRoomMessage(message)
	}

}

func (c *Client) handleJoinRoomMessage(m Message) {
	roomID := m.Message //message from front-end

	c.joinRoom(roomID, nil)

}

func (c *Client) handleLeaveRoomMessage(m Message) {
	room := c.server.findRoomByID(m.Message)
	if room == nil {
		return
	}

	if _, ok := c.rooms[room]; ok {
		delete(c.rooms, room)
	}

	room.unregister <- c
}

func (c *Client) GetName() string {
	return c.name
}

func (c *Client) GetID() string {
	return c.ID.String()
}

func (c *Client) handleJoinPrivateRoomMessage(m Message) {

	target := c.server.findClientByID(m.Message)
	if target == nil {
		return
	}

	roomID := uuid.New().String()

	c.joinRoom(roomID, target)
	target.joinRoom(roomID, c)

}

func (c *Client) joinRoom(roomID string, sender *Client) {
	roomKind := PublicRoomKind
	if sender != nil {
		roomKind = PrivateRoomKind
	}

	room := c.server.findRoomByID(roomID)
	if room == nil {
		return
	}

	if sender == nil && roomKind == PrivateRoomKind {
		return
	}

	if !c.isInRoom(room) {
		c.rooms[room] = true
		room.register <- c
		c.notifyRoomJoined(room, sender)
	}
}

func (c *Client) isInRoom(room *Room) bool {
	_, ok := c.rooms[room]

	return ok
}

func (c *Client) notifyRoomJoined(room *Room, sender *Client) {

	message := &Message{
		Action: RoomJoinedAction,
		Target: room,
		Sender: sender,
	}

	c.send <- message.encode()

}
