package message

import (
	"encoding/json"
	"log"
)

func (m *Message) encode() []byte {
	json, err := json.Marshal(m)
	if err != nil {
		log.Printf("marshal json error: %v", err)
	}

	return json

}