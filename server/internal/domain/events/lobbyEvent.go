package events

type ELobbyEvent string

const (
	Chat       ELobbyEvent = "chat"
	Connection ELobbyEvent = "connection"
	Action     ELobbyEvent = "action"
)

type LobbyEvent struct {
	Type   ELobbyEvent
	Sender Sender
	ChatEvent
	ConnectionEvent
}

type Sender struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

type ChatEvent struct {
	Message string
}
type ConnectionEvent struct {
	IsConnected bool
}
