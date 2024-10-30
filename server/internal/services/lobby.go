package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/domain/events"
)

type Lobby struct {
	Id       string
	Name     string
	Lobby    *domain.Lobby
	Clients  map[*Client]bool
	hub      *Hub
	eventsCh chan *events.LobbyEvent
}

func NewLobby(id string, name string, hub *Hub) *Lobby {
	return &Lobby{
		Id:       id,
		Name:     name,
		Lobby:    domain.NewLobby(id, name),
		Clients:  make(map[*Client]bool),
		hub:      hub,
		eventsCh: make(chan *events.LobbyEvent, 10),
	}
}

func (l *Lobby) AddClient(client *Client) {
	l.Clients[client] = true
	event := events.LobbyEvent{
		Type:            events.Connection,
		Sender:          events.Sender{UserId: client.User.ID, Username: client.User.Username},
		ConnectionEvent: events.ConnectionEvent{IsConnected: true},
	}
	l.eventsCh <- &event
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.Clients, client)
	event := events.LobbyEvent{
		Type:            events.Connection,
		Sender:          events.Sender{UserId: client.User.ID, Username: client.User.Username},
		ConnectionEvent: events.ConnectionEvent{IsConnected: false},
	}
	l.eventsCh <- &event
}

func (l *Lobby) Run() {
	for {
		select {
		case event := <-l.eventsCh:
			fmt.Printf("event: %v\n", event)
			for client := range l.Clients {
				client.EventsCh <- event
			}
		}
	}
}
