package dto

import "go-kozel/internal/domain"

type ELobbyEvent int

const (
	EventConnection  ELobbyEvent = iota
	EventChat        ELobbyEvent = iota
	EventMoveSlot    ELobbyEvent = iota
	EventUpdate      ELobbyEvent = iota
	EventError       ELobbyEvent = iota
	EventGameState   ELobbyEvent = iota
	EventStageChange ELobbyEvent = iota
	EventNewTrump    ELobbyEvent = iota
	EventChangeStep  ELobbyEvent = iota
	EventCardAction  ELobbyEvent = iota
	EventStakeResult ELobbyEvent = iota
	EventRoundResult ELobbyEvent = iota
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

type StageChangeEvent struct {
	Type  ELobbyEvent
	Stage domain.EStage
}

type NewTrumpEvent struct {
	Type  ELobbyEvent
	Trump domain.ESuit
}

type ChangeStepEvent struct {
	Type       ELobbyEvent
	PlayerStep *domain.Player
}

type CardActionEvent struct {
	Type     ELobbyEvent
	PlayerId string
	Card     *domain.Card
}

type StakeResultEvent struct {
	Type   ELobbyEvent
	Result *domain.StakeResult
}

type RoundResultEvent struct {
	Type   ELobbyEvent
	Result *domain.RoundResult
}
