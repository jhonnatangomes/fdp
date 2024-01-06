package main

import (
	"embed"
	"fdp/server"
	"log"
	"net/http"
)

//go:embed assets/*
var assets embed.FS

func main() {
	http.HandleFunc("/", server.ServeFile(assets, "assets/index.html"))
	http.HandleFunc("/index.css", server.ServeFile(assets, "assets/index.css"))
	http.HandleFunc("/htmx.1.9.10.min.js", server.ServeFile(assets, "assets/htmx.1.9.10.min.js"))
	http.HandleFunc("/ws", server.WebSocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
