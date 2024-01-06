package server

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	MessageType MessageType `json:"type"`
	Player      *Player     `json:"player"`
	Data        string      `json:"data"`
}

type MessageType int

const (
	Join MessageType = iota
)

func decodeMessage(message []byte) (*Message, error) {
	var decodedMessage Message
	if err := json.Unmarshal(message, &decodedMessage); err != nil {
		return nil, err
	}
	return &decodedMessage, nil
}

var joinedPlayers int

func (m Message) calculateResponse() *Message {
	switch m.MessageType {
	case Join:
		joinedPlayers++
		return &Message{
			MessageType: Join,
			Data:        fmt.Sprintf("Joined players: %d", joinedPlayers),
			Player:      m.Player,
		}
	}
	return nil
}
