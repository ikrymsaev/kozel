package game

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/hub"
	"log"
)

type HubEvent struct {
	Event   string
	LobbyId string
	Data    *domain.Lobby
}

type LobbyHub struct {
	Lobbies map[string]*Lobby
	hub     *hub.Hub
}

func NewLobbyHub(hub *hub.Hub) *LobbyHub {
	return &LobbyHub{
		Lobbies: make(map[string]*Lobby),
		hub:     hub,
	}
}

func (h *LobbyHub) GetHub() *hub.Hub {
	return h.hub
}

func (h *LobbyHub) CreateNewLobby(id string, name string) *Lobby {
	if _, ok := h.Lobbies[id]; ok {
		fmt.Printf("lobby with id %s already exists\n", id)
		return nil
	}

	newLobby := NewLobby(id, name, h)
	h.Lobbies[id] = newLobby
	log.Printf("lobby with id %s created", id)
	go newLobby.Run()

	log.Printf("lobby with id %s added to hub", id)
	fmt.Printf("lobby hub ch %v", h.hub.LobbyCh)
	h.hub.LobbyCh <- hub.LobbyMessage{
		Type: hub.EHubMessageNewLobby,
		Id:   id,
	}

	log.Printf("returning lobby %v", newLobby)
	return newLobby
}

func (h *LobbyHub) GetLobbyService(id string) *Lobby {
	if _, ok := h.Lobbies[id]; !ok {
		fmt.Printf("lobby with id %s not found\n", id)
		return nil
	}
	lobby := h.Lobbies[id]
	return lobby
}

func (h *LobbyHub) HasLobby(id string) bool {
	_, ok := h.Lobbies[id]
	return ok
}

func (h *LobbyHub) removeLobby(id string) {
	delete(h.Lobbies, id)
	h.hub.LobbyCh <- hub.LobbyMessage{
		Type: hub.EHubMessageRemoveLobby,
		Id:   id,
	}
}
