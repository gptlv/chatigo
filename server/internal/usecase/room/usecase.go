package room

import "github.com/gptlv/chatigo/server/internal/domain"

type roomUsecase struct {
	roomRepo domain.RoomRepository
}

func NewRoomUsecase(rr domain.RoomRepository) domain.RoomUsecase {
	return &roomUsecase{
		roomRepo: rr,
	}
}
