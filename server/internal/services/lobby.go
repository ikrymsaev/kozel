package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
)

type LobbyService struct {
	Id            string
	Name          string
	Lobby         *domain.Lobby
	GameService   *GameService
	Clients       map[*Client]bool
	hub           *Hub
	chatCh        chan *dto.ChatEvent
	connectionsCh chan *dto.ConnectionEvent
	updateCh      chan *dto.UpdateEvent
}

func NewLobbyService(id string, name string, hub *Hub) *LobbyService {
	return &LobbyService{
		Id:            id,
		Name:          name,
		Lobby:         domain.NewLobby(id, name),
		Clients:       make(map[*Client]bool),
		hub:           hub,
		chatCh:        make(chan *dto.ChatEvent, 1),
		connectionsCh: make(chan *dto.ConnectionEvent, 1),
		updateCh:      make(chan *dto.UpdateEvent, 1),
	}
}

func (l *LobbyService) AddClient(client *Client) error {
	if err := l.Lobby.ConnectPlayer(client.User); err != nil {
		return err
	}
	l.Clients[client] = true
	event := dto.ConnectionEvent{
		IsConnected: true,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
	l.sendUpdates()

	return nil
}

func (l *LobbyService) RemoveClient(client *Client) {
	delete(l.Clients, client)

	if err := l.Lobby.DisconnectPlayer(client.User); err != nil {
		fmt.Println(err)
	}
	if len(l.Clients) == 0 {
		l.hub.RemoveLobby(l.Id)
		return
	}
	event := dto.ConnectionEvent{
		IsConnected: false,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
	l.sendUpdates()
}

func (l *LobbyService) MoveSlot(client *Client, from int, to int) {
	if client.User.ID != l.Lobby.OwnerId {
		fmt.Println("Only owner can move slot")
		return
	}
	l.Lobby.MoveSlot(from, to)
	l.sendUpdates()
}

// Старт игры
func (l *LobbyService) StartGame(cl *Client) {
	if l.GameService != nil {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "Game already started",
		}
		return
	}
	if cl.User.ID != l.Lobby.OwnerId {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "Only owner can start game",
		}
		return
	}
	gameService := NewGameService(l)
	l.GameService = &gameService
	go gameService.Run()
}

func (l *LobbyService) sendUpdates() {
	event := dto.UpdateEvent{
		Type:  dto.EventUpdate,
		Slots: l.Lobby.GetSlots(),
	}

	l.updateCh <- &event
}

func (l *LobbyService) Run() {
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
