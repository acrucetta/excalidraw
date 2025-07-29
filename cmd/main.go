package main

import (
	"log"
	hub "multi-draw/internal/hub"
	"multi-draw/internal/jsonlog"
	"net/http"
	"os"
)

func serveLobby(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "web/home.html")
}

func main() {
	h := hub.NewHub()
	go h.Run()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// Home needs to be the code UI
	// Once the user enters the code, we open a new room

	// TODO: New room functionality
	http.HandleFunc("/", serveLobby)

	// User submits code
	// http.HandleFunc("/join", serveHome)

	// Redirect to room/code
	// http.HandleFunc("/room/{code}", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(h, w, r, logger)
	})

	logger.PrintInfo("Server stating", map[string]any{"port": "8080"})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
