package services

import (
	"fmt"
	"go-kozel/internal/domain"
)

type Lobby struct {
	Id      string
	Name    string
	Lobby   *domain.Lobby
	Clients map[*Client]bool
	hub     *Hub
	msgCh   chan string
}

func NewLobby(id string, name string, hub *Hub) *Lobby {
	return &Lobby{
		Id:      id,
		Name:    name,
		Lobby:   domain.NewLobby(id, name),
		Clients: make(map[*Client]bool),
		hub:     hub,
		msgCh:   make(chan string, 10),
	}
}

func (l *Lobby) AddClient(client *Client) {
	l.Clients[client] = true
	msg := fmt.Sprintf("%s joined:", client.User.Username)
	l.msgCh <- msg
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.Clients, client)
	msg := fmt.Sprintf("%s left:", client.User.Username)
	l.msgCh <- msg
}

func (l *Lobby) Run() {
	for {
		select {
		case msg := <-l.msgCh:
			for client := range l.Clients {
				client.MessageCh <- msg
			}
		}
	}
}
