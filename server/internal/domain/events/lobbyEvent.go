package events

import "go-kozel/internal/domain"

type ELobbyEvent string

const (
	Connection ELobbyEvent = "connection"
	Chat       ELobbyEvent = "chat"
	MoveSlot   ELobbyEvent = "move_slot_action"
	Update     ELobbyEvent = "update"
	Error      ELobbyEvent = "error"
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
