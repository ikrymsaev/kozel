package services

import (
	"fmt"
	"go-kozel/internal/domain"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn      *websocket.Conn
	Lobby     *domain.Lobby
	User      *domain.User
	MessageCh chan string
}

func NewClient(lobby *domain.Lobby, user *domain.User, conn *websocket.Conn) *Client {
	return &Client{
		Conn:      conn,
		Lobby:     lobby,
		User:      user,
		MessageCh: make(chan string, 10),
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.MessageCh
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()

		if err != nil {
			c.Conn.Close()
			break
		}
		fmt.Println(string(message))
	}
}
