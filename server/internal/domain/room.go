package domain

import (
	"context"
)

type Room struct {
	Id     int64
	UserId int64
	Name   string
	Kind   string
}

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
	GetRoomById(ctx context.Context, Id string) (*Room, error)
}

type RoomUsecase interface {
	CreateRoom(c context.Context, req *Room) (*Room, error)
}
