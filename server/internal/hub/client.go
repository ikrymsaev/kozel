package hub

import (
	"go-kozel/internal/domain"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	User *domain.User
	hub  *Hub
}

func (c *Client) Listen() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case event := <-c.hub.LobbyCh:
			c.Conn.WriteJSON(event)
		}
	}
}
