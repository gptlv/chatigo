package server

func (s *Server) findUserById(Id string) *Client {
	var foundClient *Client

	for client := range s.clients {
		if client.GetId() == Id {
			foundClient = client
			break
		}
	}

	return foundClient
}
