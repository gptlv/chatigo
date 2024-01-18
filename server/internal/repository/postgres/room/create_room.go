package room

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const createRoomQuery = `
INSERT INTO rooms(name, kind) VALUES ($1, $2) returning id
`

func (rr *roomRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	var lastInsertedId int

	err := rr.db.QueryRowContext(ctx, createRoomQuery, room.Name, room.Kind).Scan(&lastInsertedId)
	if err != nil {
		return nil, err
	}

	room.Id = int64(lastInsertedId)

	return room, nil
}
