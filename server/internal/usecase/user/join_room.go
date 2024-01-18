package user

import "github.com/gptlv/chatigo/server/internal/domain"

func (c *Client) joinRoom(roomId string, sender domain.User) {
	roomKind := PublicRoom
	if sender != nil {
		roomKind = PrivateRoom
	}

	room := c.server.findRoomById(roomId)
	if room == nil {
		return
	}

	if sender == nil && roomKind == PrivateRoom {
		return
	}

	if !c.isInRoom(room) {
		c.rooms[room] = true
		room.register <- c
		c.notifyRoomJoined(room, sender)
	}
}
