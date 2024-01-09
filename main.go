package main

import (
	"embed"
	"fdp/server"
	"log"
	"net/http"
)

//go:embed assets/*
var assets embed.FS

//go:embed templates/*
var templates embed.FS

func main() {
	hub := server.NewHub(&templates)
	go hub.Run()
	http.HandleFunc("/", server.ServeFile(assets, "assets/index.html"))
	http.HandleFunc("/index.css", server.ServeFile(assets, "assets/index.css"))
	http.HandleFunc("/htmx.1.9.10.min.js", server.ServeFile(assets, "assets/htmx.1.9.10.min.js"))
	http.HandleFunc("/htmx.1.9.10.ws.js", server.ServeFile(assets, "assets/htmx.1.9.10.ws.js"))
	http.HandleFunc("/game", server.WebSocketHandler(hub))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
