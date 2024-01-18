package room

func (r *Room) sendMessageToClientsInRoom(m []byte) {
	for client := range r.clients {
		client.send <- m
	}
}
