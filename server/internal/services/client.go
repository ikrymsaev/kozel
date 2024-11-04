package services

import (
	"encoding/json"
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn          *websocket.Conn
	Lobby         *LobbyService
	User          *domain.User
	chatCh        chan *dto.ChatEvent
	connectionsCh chan *dto.ConnectionEvent
	updateCh      chan *dto.UpdateEvent
	errorCh       chan *dto.ErrorEvent
	gameStateCh   chan *dto.GameStateEvent
}

func NewClient(lobby *LobbyService, user *domain.User, conn *websocket.Conn) *Client {
	return &Client{
		Conn:          conn,
		Lobby:         lobby,
		User:          user,
		chatCh:        make(chan *dto.ChatEvent, 1),
		connectionsCh: make(chan *dto.ConnectionEvent, 1),
		updateCh:      make(chan *dto.UpdateEvent, 1),
		errorCh:       make(chan *dto.ErrorEvent, 1),
		gameStateCh:   make(chan *dto.GameStateEvent, 1),
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
		case event := <-c.gameStateCh:
			c.Conn.WriteJSON(c.getGameStateMsg(event))
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
		var wsMessage = dto.WsAction{}
		marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
		if marshalErr != nil {
			log.Printf("error: %v", err)
			break
		}
		switch wsMessage.Type {
		case dto.ESendMessage:
			c.parseSendMsgAction(recievedMessage)
		case dto.EMoveSlot:
			c.parseMoveSlotAction(recievedMessage)
		case dto.EStartGame:
			c.parseStartGameAction()
		default:
			fmt.Printf("unknown action: %v\n", wsMessage.Type)
		}
	}
}

func (c *Client) parseStartGameAction() {
	c.Lobby.StartGame(c)
}

func (c *Client) parseMoveSlotAction(recievedMessage []byte) {
	var wsMessage = dto.MoveSlotAction{}
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
	event := dto.ChatEvent{
		Type:    dto.EChatEvent,
		Message: message,
		Sender:  *c.User,
	}
	c.Lobby.chatCh <- &event
}

func (c *Client) getChatMsg(event *dto.ChatEvent) dto.ChatNewMessage {
	return dto.ChatNewMessage{
		Type:    dto.NewMessage,
		Message: event.Message,
		Sender:  event.Sender,
	}
}
func (c *Client) getConnMsg(event *dto.ConnectionEvent) dto.ConnectionMessage {
	fmt.Printf("getConnMsg: %v\n", event)
	return dto.ConnectionMessage{
		Type:        dto.Connection,
		IsConnected: event.IsConnected,
		User:        event.User,
	}
}
func (c *Client) getUpdateMsg(event *dto.UpdateEvent) dto.UpdateSlotsMessage {
	return dto.UpdateSlotsMessage{
		Type:  dto.UpdateSlots,
		Slots: event.Slots,
	}
}
func (c *Client) getErrorMsg(event *dto.ErrorEvent) dto.ErrorMessage {
	return dto.ErrorMessage{
		Type:  dto.Error,
		Error: event.Error,
	}
}
func (c *Client) getGameStateMsg(event *dto.GameStateEvent) dto.GameStateMessage {
	return dto.GameStateMessage{
		Type: dto.GameState,
		Game: dto.NewGameStateModel(&event.Game),
	}
}
