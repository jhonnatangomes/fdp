package server

import (
	"embed"
	"encoding/json"
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

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	readLoop(conn)
}

func readLoop(conn *websocket.Conn) {
	for {
		messageType, messageSlice, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message, err := decodeMessage(messageSlice)
		if err != nil {
			log.Println(err)
			return
		}
		responseMessage := message.calculateResponse()
		response, err := json.Marshal(responseMessage)
		if err != nil {
			log.Println(err)
			return
		}
		if err = conn.WriteMessage(messageType, response); err != nil {
			log.Println(err)
			return
		}
	}
}
