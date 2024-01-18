package message

import (
	"database/sql"

	"github.com/gptlv/chatigo/server/internal/domain"
)

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) domain.MessageRepository {
	return &messageRepository{db: db}
}
