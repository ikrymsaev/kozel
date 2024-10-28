package events

import "go-kozel/internal/domain"

type ELobbyEvent string

const (
	Chat       ELobbyEvent = "chat"
	Connection ELobbyEvent = "connection"
	Action     ELobbyEvent = "action"
)

type LobbyEvent struct {
	Name ELobbyEvent `json:"type"`
	ChatEvent
}

type ChatEvent struct {
	Sender  *domain.User `json:"sender"`
	Message string       `json:"message"`
}

func NewChatEvent(message string, sender *domain.User) *LobbyEvent {
	return &LobbyEvent{
		Name: Chat,
		ChatEvent: ChatEvent{
			Message: message,
			Sender:  sender,
		},
	}
}

func NewLobbyEvent(message string) *LobbyEvent {
	return &LobbyEvent{
		Name: Chat,
		ChatEvent: ChatEvent{
			Message: message,
		},
	}
}
