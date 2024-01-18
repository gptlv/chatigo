package room

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const getRoomByIdQuery = `
SELECT id, name, kind FROM rooms WHERE id = $1
`

func (rr *roomRepository) GetRoomById(ctx context.Context, Id string) (*domain.Room, error) {
	room := domain.Room{}

	err := rr.db.QueryRowContext(ctx, getRoomByIdQuery, Id).Scan(&room.Id, &room.Name, &room.Kind)
	if err != nil {
		return nil, err
	}

	return &room, nil
}
