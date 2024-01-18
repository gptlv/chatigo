package server

import (
	"log"

	"github.com/gptlv/chatigo/server/internal/domain"
)

func (s *Server) runRoomFromRepository(Id string) (*domain.Room, error) {
	var room *domain.Room

	dbRoom, err := s.roomRepository.GetRoomById(ctx, Id)
	if err != nil {
		log.Printf("Error getting room by id: %v", err)
		return nil, err
	}

	room = NewRoom(dbRoom.Name, RoomKind(dbRoom.Kind))
	room.Id = dbRoom.Id

	go room.RunRoom()
	s.rooms[room] = true

	return room, nil
}
