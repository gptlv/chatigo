package server

import "fmt"

const clientJoinedMessage = "%s just joined the chat!"

func (s *Server) notifyClientJoined(c *Client) {
	message := &Message{
		Message: fmt.Sprintf(clientJoinedMessage, c.GetName()),
		Action:  UserJoinedAction,
	}
	s.broadcastToClients(message.encode())

}
