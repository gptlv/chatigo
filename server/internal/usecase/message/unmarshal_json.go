package message

import "encoding/json"

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message

	newMessage := &struct {
		*Alias
		Sender Client
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &newMessage); err != nil {
		return err
	}

	m.Sender = &newMessage.Sender

	return nil
}
