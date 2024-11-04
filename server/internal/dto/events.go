package dto

import "go-kozel/internal/domain"

type ELobbyEvent int

const (
	EventConnection ELobbyEvent = iota
	EventChat       ELobbyEvent = iota
	EventMoveSlot   ELobbyEvent = iota
	EventUpdate     ELobbyEvent = iota
	EventError      ELobbyEvent = iota
	EventGameState  ELobbyEvent = iota
)

type ErrorEvent struct {
	Type  ELobbyEvent
	Error string
}

type ChatEvent struct {
	Type     ELobbyEvent
	IsSystem bool
	Message  string
	Sender   domain.User
}

type ConnectionEvent struct {
	Type        ELobbyEvent
	IsConnected bool
	User        domain.User
}

type UpdateEvent struct {
	Type  ELobbyEvent
	Slots [4]domain.Slot
}

type GameStateEvent struct {
	Type ELobbyEvent
	Game *domain.Game
}
