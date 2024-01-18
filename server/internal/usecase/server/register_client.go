package server

func (s *Server) registerClient(c *Client) {
	s.userRepository.CreateUser(ctx, c)

	s.notifyClientJoined(c)
	s.listOnlineClients(c)
	s.clients[c] = true

	s.users = append(s.users, c)
}
