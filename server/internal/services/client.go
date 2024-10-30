package services

import (
	"encoding/json"
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/domain/events"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn          *websocket.Conn
	Lobby         *Lobby
	User          *domain.User
	chatCh        chan *events.ChatEvent
	connectionsCh chan *events.ConnectionEvent
}

func NewClient(lobby *Lobby, user *domain.User, conn *websocket.Conn) *Client {
	return &Client{
		Conn:          conn,
		Lobby:         lobby,
		User:          user,
		chatCh:        make(chan *events.ChatEvent, 1),
		connectionsCh: make(chan *events.ConnectionEvent, 1),
	}
}

// Отправляем сообщения клиенту
func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case event := <-c.chatCh:
			c.Conn.WriteJSON(c.getChatMsg(event))
		case event := <-c.connectionsCh:
			c.Conn.WriteJSON(c.getConnMsg(event))
		}
	}
}

// Получаем сообщения от клиента
func (c *Client) ReadMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, event, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var wsMessage = map[string]interface{}{}
		marshalErr := json.Unmarshal(event, &wsMessage)
		if marshalErr != nil {
			log.Printf("error: %v", err)
			break
		}
		if wsMessage["type"] == string(events.Chat) {
			data := wsMessage["data"].(map[string]interface{})
			message := data["message"].(string)
			fmt.Printf("new message: %s\n", message)
			fmt.Printf("c.User: %v\n", c.User)
			event := events.ChatEvent{
				Type:    events.Chat,
				Message: message,
				Sender:  *c.User,
			}
			c.Lobby.chatCh <- &event
		}
	}
}

type WsMessage struct {
	Type events.ELobbyEvent `json:"type"`
	Data any                `json:"data"`
}

func (c *Client) getChatMsg(event *events.ChatEvent) ChatMessage {
	fmt.Printf("getChatMsg: %v\n", event)
	return ChatMessage{
		Type:    Chat,
		Message: event.Message,
		Sender:  event.Sender,
	}
}
func (c *Client) getConnMsg(event *events.ConnectionEvent) ConnectionMessage {
	fmt.Printf("getConnMsg: %v\n", event)
	return ConnectionMessage{
		Type:        Connection,
		IsConnected: event.IsConnected,
		User:        event.User,
	}
}
