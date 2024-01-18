package server

func (s *Server) createRoom(name string, kind RoomKind) *Room {
	room := NewRoom(name, kind)

	s.roomRepository.CreateRoom(ctx, room)

	go room.RunRoom()
	s.rooms[room] = true
	return room
}
