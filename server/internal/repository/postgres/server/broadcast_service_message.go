package server

func (sr *serverRepository) broadcastServiceMessage(message []byte) {
	for u := range sr.users {
		u.send <- message
	}
}
