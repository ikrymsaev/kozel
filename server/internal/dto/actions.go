package dto

import "go-kozel/internal/domain"

// @Actions recieved from the client
type EWSAction int

const (
	WSActionSendMessage EWSAction = iota
	WSActionMoveSlot    EWSAction = iota
	WSActionStartGame   EWSAction = iota
	WSActionPraiseTrump EWSAction = iota
	WSActionMoveCard    EWSAction = iota
)

type WsAction struct {
	Type EWSAction `json:"type"`
}

type SendMessageAction struct {
	Type    EWSMessage `json:"type"`
	Message string     `json:"message"`
}
type MoveSlotAction struct {
	Type EWSMessage `json:"type"`
	From int        `json:"from"`
	To   int        `json:"to"`
}

type StartGameAction struct {
	Type EWSMessage `json:"type"`
}

type PraiseTrumpAction struct {
	Type  EWSMessage   `json:"type"`
	Trump domain.ESuit `json:"trump"`
}

type MoveCardAction struct {
	Type   EWSMessage `json:"type"`
	CardId string     `json:"cardId"`
}
