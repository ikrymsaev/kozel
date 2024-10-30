package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/domain/events"
)

type Lobby struct {
	Id            string
	Name          string
	Lobby         *domain.Lobby
	Clients       map[*Client]bool
	hub           *Hub
	chatCh        chan *events.ChatEvent
	connectionsCh chan *events.ConnectionEvent
}

func NewLobby(id string, name string, hub *Hub) *Lobby {
	return &Lobby{
		Id:            id,
		Name:          name,
		Lobby:         domain.NewLobby(id, name),
		Clients:       make(map[*Client]bool),
		hub:           hub,
		chatCh:        make(chan *events.ChatEvent, 1),
		connectionsCh: make(chan *events.ConnectionEvent, 10),
	}
}

func (l *Lobby) AddClient(client *Client) {
	l.Clients[client] = true
	event := events.ConnectionEvent{
		IsConnected: true,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.Clients, client)
	event := events.ConnectionEvent{
		IsConnected: false,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
}

func (l *Lobby) Run() {
	for {
		select {
		case event := <-l.connectionsCh:
			fmt.Printf("Lobby connectionsCh: %v\n", event)
			for client := range l.Clients {
				client.connectionsCh <- event
			}
		case event := <-l.chatCh:
			fmt.Printf("Lobby chatCh: %v\n", event)
			for client := range l.Clients {
				client.chatCh <- event
			}
		}
	}
}
