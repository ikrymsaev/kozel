package ws

type MessageType int

const (
	TypeSystem MessageType = iota
	TypeChat   MessageType = iota
	TypeGame   MessageType = iota
)

type MessageInfo struct {
	MessageType MessageType `json:"messageType"`
	RoomId      string      `json:"roomId"`
}

type ChatMessageBody struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
type ChatMessage struct {
	Info MessageInfo     `json:"info"`
	Body ChatMessageBody `json:"body"`
}
