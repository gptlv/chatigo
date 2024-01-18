package user

import "github.com/gptlv/chatigo/server/internal/domain"

func (ur *userRepository) joinRoom(roomId string, sender domain.User) {
	roomKind := PublicRoom
	if sender != nil {
		roomKind = PrivateRoom
	}

	room := ur.server.findRoomById(roomId)
	if room == nil {
		return
	}

	if sender == nil && roomKind == PrivateRoom {
		return
	}

	if !ur.isInRoom(room) {
		ur.rooms[room] = true
		room.register <- c
		ur.notifyRoomJoined(room, sender)
	}
}
