package ws

import "fmt"

type Hub struct {
	Rooms        map[string]*Room
	RegisterCh   chan *Client
	UnregisterCh chan *Client
	ChatCh       chan *ChatMessage
}

func NewHub() *Hub {
	return &Hub{
		Rooms:        make(map[string]*Room),
		RegisterCh:   make(chan *Client),
		UnregisterCh: make(chan *Client),
		ChatCh:       make(chan *ChatMessage, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.RegisterCh:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.UnregisterCh:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.ChatCh <- &ChatMessage{
							Info: MessageInfo{MessageType: TypeChat, RoomId: cl.RoomID},
							Body: ChatMessageBody{
								Content:  fmt.Sprintf("< %s has been disconected", cl.Username),
								Username: "system",
							},
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.MessageCh)
				}
			}

		case m := <-h.ChatCh:
			if _, ok := h.Rooms[m.Info.RoomId]; ok {

				for _, cl := range h.Rooms[m.Info.RoomId].Clients {
					cl.MessageCh <- m
				}
			}
		}
	}
}
