package user

func (c *Client) disconnect() {
	c.server.unregister <- c
	close(c.send)
	c.conn.Close()
	for room := range c.rooms {
		room.unregister <- c
	}
}
