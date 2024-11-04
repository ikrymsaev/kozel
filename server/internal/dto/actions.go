package dto

// @Actions recieved from the client
type EWSAction int

const (
	ESendMessage EWSAction = iota
	EMoveSlot    EWSAction = iota
	EStartGame   EWSAction = iota
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
