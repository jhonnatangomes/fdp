package server

import (
	"embed"
	"html/template"
)

type Hub struct {
	game       *Game
	clients    map[*Client]bool
	broadcast  chan *template.Template
	register   chan *Client
	unregister chan *Client
	templates  *embed.FS
}

func NewHub(templates *embed.FS) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *template.Template),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		templates:  templates,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
