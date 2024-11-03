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
	errorCh       chan *events.ErrorEvent
}

func NewClient(lobby *Lobby, user *domain.User, conn *websocket.Conn) *Client {
	return &Client{
		Conn:          conn,
		Lobby:         lobby,
		User:          user,
		chatCh:        make(chan *events.ChatEvent, 1),
		connectionsCh: make(chan *events.ConnectionEvent, 1),
		updateCh:      make(chan *events.UpdateEvent, 1),
		errorCh:       make(chan *events.ErrorEvent, 1),
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
		case event := <-c.errorCh:
			c.Conn.WriteJSON(c.getErrorMsg(event))
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
		var wsMessage = WsAction{}
		marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
		if marshalErr != nil {
			log.Printf("error: %v", err)
			break
		}
		if wsMessage.Type == SendMessage {
			c.parseSendMsgAction(recievedMessage)
		}
		if wsMessage.Type == MoveSlot {
			c.parseMoveSlotAction(recievedMessage)
		}
	}
}

func (c *Client) parseMoveSlotAction(recievedMessage []byte) {
	var wsMessage = MoveSlotAction{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}

	c.Lobby.MoveSlot(c, wsMessage.From, wsMessage.To)
}

func (c *Client) parseSendMsgAction(recievedMessage []byte) {
	var wsMessage = map[string]interface{}{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}
	message := wsMessage["message"].(string)
	fmt.Printf("new message: %s\n", message)
	fmt.Printf("c.User: %v\n", c.User)
	event := events.ChatEvent{
		Type:    events.Chat,
		Message: message,
		Sender:  *c.User,
	}
	c.Lobby.chatCh <- &event
}

func (c *Client) getChatMsg(event *events.ChatEvent) ChatNewMessage {
	fmt.Printf("getChatMsg: %v\n", event)
	return ChatNewMessage{
		Type:    NewMessage,
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
func (c *Client) getUpdateMsg(event *events.UpdateEvent) UpdateSlotsMessage {
	fmt.Printf("getUpdateMsg: %v\n", event)
	return UpdateSlotsMessage{
		Type:  UpdateSlots,
		Slots: event.Slots,
	}
}
func (c *Client) getErrorMsg(event *events.ErrorEvent) ErrorMessage {
	fmt.Printf("getErrorMsg: %v\n", event)
	return ErrorMessage{
		Type:  Error,
		Error: event.Error,
	}
}
