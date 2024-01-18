package room

func (r *Room) unregisterClient(c *Client) {
	if _, ok := r.clients[c]; ok {
		delete(r.clients, c)
	}
}
