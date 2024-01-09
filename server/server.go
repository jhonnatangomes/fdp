package server

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeFile(assets embed.FS, filePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := assets.ReadFile(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var contentType string
		switch {
		case strings.HasSuffix(filePath, ".html"):
			contentType = "text/html"
		case strings.HasSuffix(filePath, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(filePath, ".js"):
			contentType = "text/javascript"
		}
		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	}
}

func WebSocketHandler(hub *Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{
			hub:  hub,
			conn: conn,
			send: make(chan *template.Template),
		}
		client.hub.register <- client
		go client.writePump()
		go client.readPump()
	}
}
