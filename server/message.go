package server

import (
	"html/template"
)

type Message struct {
	MessageType MessageType `json:"type"`
	Data        string      `json:"data"`
}

type MessageType string

const (
	Join MessageType = "join"
)

func (m Message) calculateResponse(hub *Hub) *template.Template {
	switch m.MessageType {
	case Join:
		if hub.game == nil {
			hub.game = newGame()
		}
		hub.game.addPlayer(m.Data)
		data, err := hub.templates.ReadFile("templates/player_joined.html")
		if err != nil {
			return nil
		}
		templ, err := template.New("player_joined").Parse(string(data))
		if err != nil {
			return nil
		}
		return templ
	}
	return nil
}
