package server

func (s *Server) listOnlineClients(c *Client) {
	for _, user := range s.users {
		message := &Message{
			Action: UserJoinedAction,
			Sender: user,
		}
		c.send <- message.encode()
	}
}
