package services

import "go-kozel/internal/domain"

type EMessageType string

const (
	Connection EMessageType = "connection"
	Chat       EMessageType = "chat"
	MoveSlot   EMessageType = "move_slot_action"
	Update     EMessageType = "update"
	Error      EMessageType = "error"
)

type WsMessage struct {
	Type EMessageType `json:"type"`
}

type ErrorMessage struct {
	Type  EMessageType `json:"type"`
	Error string       `json:"error"`
}

type ChatMessage struct {
	Type     EMessageType `json:"type"`
	Sender   domain.User  `json:"sender"`
	IsSystem bool         `json:"isSystem"`
	Message  string       `json:"message"`
}

type ConnectionMessage struct {
	Type        EMessageType `json:"type"`
	User        domain.User  `json:"user"`
	IsConnected bool         `json:"isConnected"`
}

type UpdateMessage struct {
	Type  EMessageType   `json:"type"`
	Slots [4]domain.Slot `json:"slots"`
}

type ActionMessage struct {
	Type   EMessageType `json:"type"`
	Action string       `json:"action"`
}
type MoveSlotMessage struct {
	Type EMessageType `json:"type"`
	From int          `json:"from"`
	To   int          `json:"to"`
}
