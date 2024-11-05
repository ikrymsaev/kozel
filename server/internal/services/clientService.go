package services

import (
	"encoding/json"
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
	"log"

	"github.com/gorilla/websocket"
)

type ClientService struct {
	Conn          *websocket.Conn
	LobbyService  *LobbyService
	User          *domain.User
	chatCh        chan *dto.ChatEvent
	connectionsCh chan *dto.ConnectionEvent
	updateCh      chan *dto.UpdateEvent
	errorCh       chan *dto.ErrorEvent
	gameStateCh   chan *dto.GameStateEvent
	stageCh       chan *dto.StageChangeEvent
	trumpCh       chan *dto.NewTrumpEvent
	playerStepCh  chan *dto.ChangeStepEvent
	cardActionCh  chan *dto.CardActionEvent
}

func NewClientService(lobby *LobbyService, user *domain.User, conn *websocket.Conn) *ClientService {
	return &ClientService{
		Conn:          conn,
		LobbyService:  lobby,
		User:          user,
		chatCh:        make(chan *dto.ChatEvent, 1),
		connectionsCh: make(chan *dto.ConnectionEvent, 1),
		updateCh:      make(chan *dto.UpdateEvent, 1),
		errorCh:       make(chan *dto.ErrorEvent, 1),
		gameStateCh:   make(chan *dto.GameStateEvent, 1),
		stageCh:       make(chan *dto.StageChangeEvent),
		trumpCh:       make(chan *dto.NewTrumpEvent, 1),
		playerStepCh:  make(chan *dto.ChangeStepEvent, 1),
		cardActionCh:  make(chan *dto.CardActionEvent, 1),
	}
}

// Отправляем сообщения клиенту
func (c *ClientService) WriteMessage() {
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
		case event := <-c.stageCh:
			c.Conn.WriteJSON(c.getStageMsg(event))
		case event := <-c.trumpCh:
			c.Conn.WriteJSON(c.getTrumpMsg(event))
		case event := <-c.playerStepCh:
			c.Conn.WriteJSON(c.getChangeStepMsg(event))
		case event := <-c.cardActionCh:
			c.Conn.WriteJSON(c.getCardActionMsg(event))
		}
	}
}

// Получаем сообщения от клиента
func (c *ClientService) ReadMessage() {
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
		case dto.WSActionSendMessage:
			c.parseSendMsgAction(recievedMessage)
		case dto.WSActionMoveSlot:
			c.parseMoveSlotAction(recievedMessage)
		case dto.WSActionStartGame:
			c.parseStartGameAction()
		case dto.WSActionPraiseTrump:
			c.parsePraiseTrumpAction(recievedMessage)
		case dto.WSActionMoveCard:
			c.parseMoveCardAction(recievedMessage)
		default:
			fmt.Printf("unknown action: %v\n", wsMessage.Type)
		}
	}
}

func (c *ClientService) parseMoveCardAction(recievedMessage []byte) {
	var wsMessage = dto.MoveCardAction{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}
	c.LobbyService.GameService.MoveCard(c, wsMessage.CardId)
}

func (c *ClientService) parseStartGameAction() {
	c.LobbyService.StartGame(c)
}

func (c *ClientService) parsePraiseTrumpAction(recievedMessage []byte) {
	var wsMessage = dto.PraiseTrumpAction{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}
	c.LobbyService.GameService.PraiseTrump(c, &wsMessage.Trump)
}

func (c *ClientService) parseMoveSlotAction(recievedMessage []byte) {
	var wsMessage = dto.MoveSlotAction{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}

	c.LobbyService.MoveSlot(c, wsMessage.From, wsMessage.To)
}

func (c *ClientService) parseSendMsgAction(recievedMessage []byte) {
	var wsMessage = map[string]interface{}{}
	marshalErr := json.Unmarshal(recievedMessage, &wsMessage)
	if marshalErr != nil {
		log.Printf("error: %v", marshalErr)
		return
	}
	message := wsMessage["message"].(string)
	event := dto.ChatEvent{
		Type:    dto.EventChat,
		Message: message,
		Sender:  *c.User,
	}
	c.LobbyService.chatCh <- &event
}

func (c *ClientService) getChatMsg(event *dto.ChatEvent) dto.ChatNewMessage {
	return dto.ChatNewMessage{
		Type:    dto.WSMessageNewMessage,
		Message: event.Message,
		Sender:  event.Sender,
	}
}
func (c *ClientService) getConnMsg(event *dto.ConnectionEvent) dto.ConnectionMessage {
	fmt.Printf("getConnMsg: %v\n", event)
	return dto.ConnectionMessage{
		Type:        dto.WSMessageConnection,
		IsConnected: event.IsConnected,
		User:        event.User,
	}
}
func (c *ClientService) getUpdateMsg(event *dto.UpdateEvent) dto.UpdateSlotsMessage {
	return dto.UpdateSlotsMessage{
		Type:  dto.WSMessageUpdateSlots,
		Slots: event.Slots,
	}
}
func (c *ClientService) getErrorMsg(event *dto.ErrorEvent) dto.ErrorMessage {
	return dto.ErrorMessage{
		Type:  dto.WSMessageError,
		Error: event.Error,
	}
}
func (c *ClientService) getGameStateMsg(event *dto.GameStateEvent) dto.GameStateMessage {
	gameModel := dto.NewGameStateModel(event.Game)

	for index, player := range gameModel.Players {
		if player.User == nil || player.User.ID != c.User.ID {
			hiddenHand := []dto.CardStateModel{}
			for _, card := range player.Hand {
				hiddenHand = append(hiddenHand, dto.CardStateModel{Id: card.Id, IsHidden: true})
			}
			gameModel.Players[index].Hand = hiddenHand
		}
	}

	return dto.GameStateMessage{
		Type: dto.WSMessageGameState,
		Game: gameModel,
	}
}
func (c *ClientService) getStageMsg(event *dto.StageChangeEvent) dto.StageMessage {
	return dto.StageMessage{
		Type:  dto.WSMessageStage,
		Stage: event.Stage,
	}
}
func (c *ClientService) getTrumpMsg(event *dto.NewTrumpEvent) dto.NewTrumpMessage {
	return dto.NewTrumpMessage{
		Type:  dto.WSMEssageNewTrump,
		Trump: event.Trump,
	}
}

func (c *ClientService) getChangeStepMsg(event *dto.ChangeStepEvent) dto.ChangeTurnMessage {
	return dto.ChangeTurnMessage{
		Type:         dto.WSMessageChangeTurn,
		TurnPlayerId: event.PlayerStep.Id,
	}
}

func (c *ClientService) getCardActionMsg(event *dto.CardActionEvent) dto.CardActionMessage {
	return dto.CardActionMessage{
		Type:     dto.WSMessageCardAction,
		PlayerId: event.PlayerId,
		Card:     dto.GetCardStateModel(event.Card),
	}
}
