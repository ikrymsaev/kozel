package services

import (
	"fmt"
	"go-kozel/internal/domain"
)

type HubEvent struct {
	Event   string
	LobbyId string
	Data    *domain.Lobby
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

func (h *Hub) CreateNewLobby(id string, name string) *Lobby {
	if _, ok := h.Lobbies[id]; ok {
		fmt.Printf("lobby with id %s already exists\n", id)
		return nil
	}

	newLobby := NewLobby(id, name, h)
	h.Lobbies[id] = newLobby
	go newLobby.Run()

	// Notify about creating lobby
	for listener := range h.listeners {
		*listener <- &HubEvent{Event: "new_lobby", Data: newLobby.Lobby, LobbyId: id}
	}

	return newLobby
}

func (h *Hub) GetHubLobby(id string) *Lobby {
	if _, ok := h.Lobbies[id]; !ok {
		fmt.Printf("lobby with id %s not found\n", id)
		return nil
	}
	lobby := h.Lobbies[id]
	return lobby
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
