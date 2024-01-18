package domain

import "context"

type ChatMessage struct {
	Id      string
	Context string
	// Timestamp
	SenderId string
	RoomId   string
}

type ChatMessageRepository interface {
	CreateMessage(ctx context.Context, chatMessage *ChatMessage) (*ChatMessage, error)
	DeleteMessage(ctx context.Context, chatMessage *ChatMessage) error
	GetMessageById(id string) (*ChatMessage, error)
}
