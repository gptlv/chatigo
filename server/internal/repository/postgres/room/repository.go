package room

import (
	"database/sql"

	"github.com/gptlv/chatigo/server/internal/domain"
)

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) domain.RoomRepository {
	return &roomRepository{db: db}
}
