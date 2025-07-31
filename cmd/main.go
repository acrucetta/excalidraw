package main

import (
	"log"
	"multi-draw/internal/canvas"
	"multi-draw/internal/hub"
	"multi-draw/internal/jsonlog"
	"multi-draw/internal/rooms"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Room Helpers

var globalRooms = make(map[int]*rooms.Room)
var roomsMu sync.Mutex

func getOrCreateRoom(roomCode int) *rooms.Room {
	roomsMu.Lock()
	defer roomsMu.Unlock()
	room, ok := globalRooms[roomCode]
	if !ok {
		log.Printf("Creating new room: %d", roomCode)
		room = rooms.NewRoom(roomCode)
		go room.Hub.Run() // start hub for this room
		globalRooms[roomCode] = room
	} else {
		log.Printf("Using existing room: %d", roomCode)
	}
	return room
}

// Handlers

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(w http.ResponseWriter, r *http.Request, logger *jsonlog.Logger) {
	roomCode := r.URL.Query().Get("room")
	if roomCode == "" {
		http.Error(w, "Room code required", http.StatusBadRequest)
		return
	}
	roomCodeInt, err := strconv.Atoi(roomCode)
	if err != nil {
		http.Error(w, "Invalid room code", http.StatusBadRequest)
		return
	}
	room := getOrCreateRoom(roomCodeInt)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.PrintError(err, nil)
		return
	}

	client := &hub.Client{
		Hub:    room.Hub,
		Conn:   conn,
		Send:   make(chan canvas.StrokeSegment, 256),
		Logger: logger,
	}
	for _, stroke := range room.Strokes {
		client.Send <- stroke
	}
	room.Hub.Register <- client
	go client.WritePump()
	go client.ReadPump()
}

func serveLobby(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "web/lobby.html")
}

func serveRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	code := vars["code"]
	log.Println("Serving room:", code)

	// Serve the room HTML. If you want to pass the code, see note below.
	http.ServeFile(w, r, "web/room.html")
}

func main() {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	r := mux.NewRouter()
	r.HandleFunc("/", serveLobby)
	r.HandleFunc("/room/{code}", serveRoom)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r, logger)
	})

	logger.PrintInfo("Server stating", map[string]any{"port": "8080"})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
