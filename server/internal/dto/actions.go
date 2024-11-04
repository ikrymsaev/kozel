package dto

// @Actions recieved from the client
type EWSAction int

const (
	WSActionSendMessage EWSAction = iota
	WSActionMoveSlot    EWSAction = iota
	WSActionStartGame   EWSAction = iota
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
