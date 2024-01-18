package user

import "github.com/gptlv/chatigo/server/internal/domain"

func (ur *userRepository) isInRoom(room *domain.Room) bool {
	_, ok := ur.rooms[room]

	return ok
}
