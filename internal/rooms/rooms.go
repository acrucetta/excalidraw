package rooms

import (
	"fmt"
	"multi-draw/internal/canvas"
	"multi-draw/internal/hub"
	"slices"

	"github.com/google/uuid"
)

type PlayerAlreadyExistsError struct {
	ItemID int
}
type PlayerNotFoundError struct {
	ItemID int
}

func (e *PlayerAlreadyExistsError) Error() string {
	return fmt.Sprintf("Item %d already exists in the list", e.ItemID)
}

func (e *PlayerNotFoundError) Error() string {
	return fmt.Sprintf("Item %d not found in the list", e.ItemID)
}

type Room struct {
	RoomID    uuid.UUID
	RoomCode  int
	PlayerIDs []int
	Hub       *hub.Hub
	Strokes   []canvas.StrokeSegment
}

func NewRoom(code int) *Room {
	room := &Room{
		RoomID:    uuid.New(),
		RoomCode:  code,
		PlayerIDs: make([]int, 0),
		Hub:       hub.NewHub(),
		Strokes:   make([]canvas.StrokeSegment, 0),
	}
	h := hub.NewHub()
	h.History = &room.Strokes
	room.Hub = h
	return room
}

func (r *Room) AddPlayer(playerId int) (bool, error) {
	// Check if it's already in the list
	if slices.Contains(r.PlayerIDs, playerId) {
		return false, &PlayerAlreadyExistsError{ItemID: playerId}
	}
	r.PlayerIDs = append(r.PlayerIDs, playerId)
	return true, nil
}

func (r *Room) RemovePlayer(playerId int) (bool, error) {
	for i, v := range r.PlayerIDs {
		if v == playerId {
			r.PlayerIDs = append(r.PlayerIDs[:i], r.PlayerIDs[i+1:]...)
			return true, nil
		}
	}
	return false, &PlayerNotFoundError{ItemID: playerId}
}
