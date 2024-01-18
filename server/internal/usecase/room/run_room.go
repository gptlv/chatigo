package room

func (r *Room) RunRoom() {
	for {
		select {
		case client := <-r.register:
			r.registerClient(client)
		case client := <-r.unregister:
			r.unregisterClient(client)
		case message := <-r.broadcast:
			r.sendMessageToClientsInRoom(message.encode())
		}

	}
}
