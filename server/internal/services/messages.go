package services

type ChatMessage struct {
	Message string `json:"message"`
}

type ConnectionMessage struct {
	IsConnected bool `json:"isConnected"`
}
