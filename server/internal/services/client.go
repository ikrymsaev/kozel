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
	Conn     *websocket.Conn
	Lobby    *Lobby
	User     *domain.User
	EventsCh chan *events.LobbyEvent
}

func NewClient(lobby *Lobby, user *domain.User, conn *websocket.Conn) *Client {
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
		switch event.Type {
		case events.Connection:
			c.Conn.WriteJSON(c.getConnMsg(event))
		case events.Chat:
			msg := c.getChatMsg(event)
			fmt.Printf("client msg: %v\n", msg)
			c.Conn.WriteJSON(msg)
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
			event := events.LobbyEvent{
				Type:      events.Chat,
				Sender:    events.Sender{UserId: c.User.ID, Username: c.User.Username},
				ChatEvent: events.ChatEvent{Message: message},
			}
			c.Lobby.eventsCh <- &event
		}
	}
}

type WsMessage struct {
	Type   events.ELobbyEvent `json:"type"`
	Sender events.Sender      `json:"sender"`
	Data   any                `json:"data"`
}

func (c *Client) getChatMsg(event *events.LobbyEvent) WsMessage {
	fmt.Printf("getChatMsg: %v\n", event)
	return WsMessage{
		Type:   events.Chat,
		Sender: event.Sender,
		Data:   ChatMessage{Message: event.Message},
	}
}
func (c *Client) getConnMsg(event *events.LobbyEvent) WsMessage {
	return WsMessage{
		Type:   events.Connection,
		Sender: event.Sender,
		Data:   ConnectionMessage{IsConnected: event.IsConnected},
	}
}
