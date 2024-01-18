package user

func (c *Client) handleJoinRoomMessage(m Message) {
	roomID := m.Message //message from front-end

	c.joinRoom(roomID, nil)

}
