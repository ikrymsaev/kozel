package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn      *websocket.Conn
	MessageCh chan *ChatMessage
	ID        string `json:"id"`
	RoomID    string `json:"roomId"`
	Username  string `json:"username"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.MessageCh
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.UnregisterCh <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msgContent map[string]string
		marshalErr := json.Unmarshal(m, &msgContent)
		if marshalErr != nil {
			log.Printf("error: %v", err)
			break
		}

		msg := &ChatMessage{
			Info: MessageInfo{MessageType: TypeChat, RoomId: c.RoomID},
			Body: ChatMessageBody{Content: msgContent["content"], Username: c.Username},
		}
		hub.ChatCh <- msg
	}
}
