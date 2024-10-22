package services

type HubEvent struct {
	Event   string
	LobbyId string
	Data    *Lobby
}

type Hub struct {
	Lobbies   map[string]*Lobby
	listeners map[*chan *HubEvent]bool
}

func NewHub() *Hub {
	return &Hub{
		Lobbies:   make(map[string]*Lobby),
		listeners: make(map[*chan *HubEvent]bool),
	}
}

func (h *Hub) Listen(listener *chan *HubEvent) {
	h.listeners[listener] = true
}
func (h *Hub) Unlisten(listener *chan *HubEvent) {
	delete(h.listeners, listener)
}

func (h *Hub) AddLobby(lobby *Lobby) {
	h.Lobbies[lobby.OwnerId] = lobby

	for listener := range h.listeners {
		*listener <- &HubEvent{Event: "new_lobby", Data: lobby, LobbyId: lobby.Id}
	}
}

func (h *Hub) HasLobby(id string) bool {
	_, ok := h.Lobbies[id]
	return ok
}

func (h *Hub) RemoveLobby(id string) {
	delete(h.Lobbies, id)

	for listener := range h.listeners {
		*listener <- &HubEvent{Event: "remove_lobby", Data: nil, LobbyId: id}
	}
}
