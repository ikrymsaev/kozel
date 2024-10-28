package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/domain/events"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Lobby    *domain.Lobby
	User     *domain.User
	EventsCh chan *events.LobbyEvent
}

func NewClient(lobby *domain.Lobby, user *domain.User, conn *websocket.Conn) *Client {
	return &Client{
		Conn:     conn,
		Lobby:    lobby,
		User:     user,
		EventsCh: make(chan *events.LobbyEvent, 10),
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		event, ok := <-c.EventsCh
		if !ok {
			return
		}
		sender := c.User.Username
		if event.Sender != nil && event.Sender.ID == c.User.ID {
			sender = "you"
		}
		msg := fmt.Sprintf("%s: %s", sender, event.Message)
		c.Conn.WriteJSON(msg)
	}
}

func (c *Client) ReadMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, event, err := c.Conn.ReadMessage()

		if err != nil {
			c.Conn.Close()
			break
		}
		fmt.Println(string(event))
	}
}
