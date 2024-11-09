package hub

import (
	"go-kozel/internal/domain"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]*Client
	LobbyCh chan LobbyMessage
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]*Client),
		LobbyCh: make(chan LobbyMessage),
	}
}

func (h *Hub) Connect(conn *websocket.Conn, user *domain.User) {
	client := &Client{Conn: conn, User: user, hub: h}
	h.clients[conn] = client
	client.Listen()
}

func (h *Hub) Disconnect(conn *websocket.Conn) {
	delete(h.clients, conn)
}

func (h *Hub) Run() {
	for {
		select {
		case message := <-h.LobbyCh:
			for _, client := range h.clients {
				client.Conn.WriteJSON(message)
			}
		}
	}
}
