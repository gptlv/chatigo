package user

import "github.com/gptlv/chatigo/server/internal/domain"

func (c *Client) notifyRoomJoined(room *Room, sender domain.User) {

	message := &Message{
		Action: RoomJoinedAction,
		Target: room,
		Sender: sender,
	}

	c.send <- message.encode()

}
