package services

import (
	"go-kozel/internal/domain"
	"go-kozel/internal/domain/events"
)

type Lobby struct {
	Id      string
	Name    string
	Lobby   *domain.Lobby
	Clients map[*Client]bool
	hub     *Hub
	msgCh   chan *events.LobbyEvent
}

func NewLobby(id string, name string, hub *Hub) *Lobby {
	return &Lobby{
		Id:      id,
		Name:    name,
		Lobby:   domain.NewLobby(id, name),
		Clients: make(map[*Client]bool),
		hub:     hub,
		msgCh:   make(chan *events.LobbyEvent, 10),
	}
}

func (l *Lobby) AddClient(client *Client) {
	l.Clients[client] = true
	event := events.NewChatEvent("joined to lobby", client.User)
	l.msgCh <- event
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.Clients, client)
	event := events.NewChatEvent("has left from lobby", client.User)
	l.msgCh <- event
}

func (l *Lobby) Run() {
	for {
		select {
		case msg := <-l.msgCh:
			for client := range l.Clients {
				client.EventsCh <- msg
			}
		}
	}
}
