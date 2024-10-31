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
	updateCh      chan *events.UpdateEvent
}

func NewClient(lobby *Lobby, user *domain.User, conn *websocket.Conn) *Client {
	return &Client{
		Conn:          conn,
		Lobby:         lobby,
		User:          user,
		chatCh:        make(chan *events.ChatEvent, 1),
		connectionsCh: make(chan *events.ConnectionEvent, 1),
		updateCh:      make(chan *events.UpdateEvent, 1),
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
		case event := <-c.updateCh:
			c.Conn.WriteJSON(c.getUpdateMsg(event))
		}
	}
}

// Получаем сообщения от клиента
func (c *Client) ReadMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, recievedMessage, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var wsMessage = WsMessage{}
		marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
		if marshalErr != nil {
			log.Printf("error: %v", err)
			break
		}
		if string(wsMessage.Type) == string(events.Chat) {
			c.parseChatMessage(recievedMessage)
		}
		if string(wsMessage.Type) == string(events.MoveSlot) {
			c.parseMoveSlotMessage(recievedMessage)
		}
	}
}

func (c *Client) parseMoveSlotMessage(recievedMessage []byte) {
	var wsMessage = MoveSlotMessage{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}

	c.Lobby.MoveSlot(c, wsMessage.From, wsMessage.To)
}

func (c *Client) parseChatMessage(recievedMessage []byte) {
	var wsMessage = map[string]interface{}{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}
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
func (c *Client) getUpdateMsg(event *events.UpdateEvent) UpdateMessage {
	fmt.Printf("getUpdateMsg: %v\n", event)
	return UpdateMessage{
		Type:  Update,
		Slots: event.Slots,
	}
}
