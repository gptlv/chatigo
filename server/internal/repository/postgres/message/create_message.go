package message

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const createMessageQuery = `
INSERT INTO messages(content, timestamp, sender, room) VALUES ($1, $2, $3, $4) returning id
`

func (mr *messageRepository) CreateMessage(context.Context, *domain.ChatMessage) (*domain.ChatMessage, error) {

}
