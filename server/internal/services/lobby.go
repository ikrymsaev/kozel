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
	updateCh      chan *events.UpdateEvent
}

func NewLobby(id string, name string, hub *Hub) *Lobby {
	return &Lobby{
		Id:            id,
		Name:          name,
		Lobby:         domain.NewLobby(id, name),
		Clients:       make(map[*Client]bool),
		hub:           hub,
		chatCh:        make(chan *events.ChatEvent, 1),
		connectionsCh: make(chan *events.ConnectionEvent, 1),
		updateCh:      make(chan *events.UpdateEvent, 1),
	}
}

func (l *Lobby) AddClient(client *Client) {
	if err := l.Lobby.ConnectPlayer(client.User); err != nil {
		fmt.Println(err)
		return
	}
	l.Clients[client] = true
	event := events.ConnectionEvent{
		IsConnected: true,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
	l.sendUpdates()
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.Clients, client)

	if err := l.Lobby.DisconnectPlayer(client.User); err != nil {
		fmt.Println(err)
	}
	event := events.ConnectionEvent{
		IsConnected: false,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
	l.sendUpdates()
}

func (l *Lobby) MoveSlot(client *Client, from int, to int) {
	if client.User.ID != l.Lobby.OwnerId {
		fmt.Println("Only owner can move slot")
		return
	}
	l.Lobby.MoveSlot(from, to)
	l.sendUpdates()
}

func (l *Lobby) sendUpdates() {
	event := events.UpdateEvent{
		Type:  events.Update,
		Slots: l.Lobby.GetSlots(),
	}

	l.updateCh <- &event
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
		case event := <-l.updateCh:
			fmt.Printf("Lobby updateCh: %v\n", event)
			for client := range l.Clients {
				client.updateCh <- event
			}
		}
	}
}
