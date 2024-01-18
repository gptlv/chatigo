package user

import "github.com/google/uuid"

func (c *Client) handleJoinPrivateRoomMessage(m Message) {

	target := c.server.findClientById(m.Message)
	if target == nil {
		return
	}

	roomId := uuid.New().String()

	c.joinRoom(roomId, target)
	target.joinRoom(roomID, c)

}
