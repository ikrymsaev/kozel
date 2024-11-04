package dto

import "go-kozel/internal/domain"

// @Messages send to the client
type EWSMessage int

const (
	Error       EWSMessage = iota
	Connection  EWSMessage = iota
	NewMessage  EWSMessage = iota
	UpdateSlots EWSMessage = iota
	GameState   EWSMessage = iota
)

type WsMessage struct {
	Type EWSMessage `json:"type"`
}

type ErrorMessage struct {
	Type  EWSMessage `json:"type"`
	Error string     `json:"error"`
}

type ChatNewMessage struct {
	Type     EWSMessage  `json:"type"`
	Sender   domain.User `json:"sender"`
	Message  string      `json:"message"`
	IsSystem bool        `json:"isSystem"`
}

type ConnectionMessage struct {
	Type        EWSMessage  `json:"type"`
	IsConnected bool        `json:"isConnected"`
	User        domain.User `json:"user"`
}

type UpdateSlotsMessage struct {
	Type  EWSMessage     `json:"type"`
	Slots [4]domain.Slot `json:"slots"`
}

type GameStateMessage struct {
	Type EWSMessage     `json:"type"`
	Game GameStateModel `json:"game"`
}
