package user

func (c *Client) handleLeaveRoomMessage(m Message) {
	room := c.server.findRoomById(m.Message)
	if room == nil {
		return
	}

	if _, ok := c.rooms[room]; ok {
		delete(c.rooms, room)
	}

	room.unregister <- c
}
