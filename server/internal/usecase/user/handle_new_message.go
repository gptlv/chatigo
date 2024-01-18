package user

import (
	"encoding/json"
	"log"
)

func (c *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message: %v", err)
	}

	message.Sender = c

	switch message.Action {
	case SendMessageAction:
		room := message.Target

		if room := c.server.findRoomById(room.GetId()); room != nil {
			room.broadcast <- &message
		}

	case JoinRoomAction:
		c.handleJoinRoomMessage(message)

	case LeaveRoomAction:
		c.handleLeaveRoomMessage(message)

	case JoinPrivateRoomAction:
		c.handleJoinPrivateRoomMessage(message)
	}

}
