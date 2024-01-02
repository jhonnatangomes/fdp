package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	readLoop(conn)
}

var joinedPlayers int

func readLoop(conn *websocket.Conn) {
	for {
		messageType, messageSlice, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := string(messageSlice)
		if message == "join" {
			joinedPlayers++
		}
		response := []byte(fmt.Sprintf("Joined players: %d", joinedPlayers))
		if err = conn.WriteMessage(messageType, response); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
