package message

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const getMessageByIdQuery = `
SELECT content, timestamp, sender, room FROM messages WHERE id = $1
`

func (mr *messageRepository) GetMessageById(ctx context.Context, id string) (*domain.Message, error) {

}
