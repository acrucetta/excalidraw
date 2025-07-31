package hub

import canvas "multi-draw/internal/canvas"

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan canvas.StrokeSegment
	Register   chan *Client
	Unregister chan *Client
	History    *[]canvas.StrokeSegment
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan canvas.StrokeSegment, 256),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		History:    nil,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
