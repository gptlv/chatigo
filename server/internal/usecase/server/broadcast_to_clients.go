package server

func (s *Server) broadcastToClients(message []byte) {
	for c := range s.clients {
		c.send <- message
	}
}
