package rooms

import (
	"fmt"
	hub "multi-draw/internal/hub"

	"github.com/google/uuid"
)

type AlreadyExistsError struct {
	ItemID string
}
type NotFoundError struct {
	ItemID string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("Item %v already exists in the list", e.ItemID)
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Item %v not found in the list", e.ItemID)
}

type Room struct {
	RoomID    string
	RoomCode  string
	PlayerIDs []string
	Hub       *hub.Hub
}

func NewRoom(code string) *Room {
	return &Room{
		RoomID:    uuid.New().String(),
		RoomCode:  code,
		PlayerIDs: make([]string, 0),
		Hub:       hub.NewHub(),
	}
}

func (r *Room) AddPlayer(playerId string) (bool, error) {
	// Check if it's already in the list
	for _, v := range r.PlayerIDs {
		if v == playerId {
			return false, &AlreadyExistsError{ItemID: playerId}
		}
	}
	r.PlayerIDs = append(r.PlayerIDs, playerId)
	return true, nil
}

func (r *Room) RemovePlayer(playerId string) (bool, error) {
	for i, v := range r.PlayerIDs {
		if v == playerId {
			r.PlayerIDs = append(r.PlayerIDs[:i], r.PlayerIDs[i+1:]...)
			return true, nil
		}
	}
	return false, &NotFoundError{ItemID: playerId}
}
