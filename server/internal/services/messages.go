package services

import "go-kozel/internal/domain"

type EMessageType string

const (
	Connection EMessageType = "connection"
	Chat       EMessageType = "chat"
	Action     EMessageType = "action"
	Update     EMessageType = "update"
)

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
