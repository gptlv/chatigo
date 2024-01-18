package server

import "fmt"

const clientLeftMessage = "%s left"

func (s *Server) notifyClientLeft(c *Client) {
	message := &Message{
		Message: fmt.Sprintf(clientJoinedMessage, c.GetName()),
		Action:  UserLeftAction,
	}
	s.broadcastToClients(message.encode())

}
