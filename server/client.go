package server

import (
	"encoding/json"
	"html/template"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *template.Template
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var jsonMessage Message
		json.Unmarshal(message, &jsonMessage)
		template := jsonMessage.calculateResponse(c.hub)
		c.hub.broadcast <- template
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case template, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			err = template.Execute(w, c.hub.game)
			if err != nil {
				log.Println(err)
				return
			}
			if err := w.Close(); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
