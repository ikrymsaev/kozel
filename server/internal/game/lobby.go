package game

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
)

type Lobby struct {
	Id            string
	Name          string
	Lobby         *domain.Lobby
	GameService   *Game
	Clients       map[*WsClient]bool
	hub           *LobbyHub
	chatCh        chan *dto.ChatEvent
	connectionsCh chan *dto.ConnectionEvent
	updateCh      chan *dto.UpdateEvent
}

func NewLobby(id string, name string, hub *LobbyHub) *Lobby {
	return &Lobby{
		Id:            id,
		Name:          name,
		Lobby:         domain.NewLobby(id, name),
		Clients:       make(map[*WsClient]bool),
		hub:           hub,
		chatCh:        make(chan *dto.ChatEvent, 1),
		connectionsCh: make(chan *dto.ConnectionEvent, 1),
		updateCh:      make(chan *dto.UpdateEvent, 1),
	}
}

func (l *Lobby) AddClient(client *WsClient) error {
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
	if l.GameService != nil {
		client.gameStateCh <- &dto.GameStateEvent{
			Type: dto.EventGameState,
			Game: l.GameService.Game,
		}
	}

	return nil
}

func (l *Lobby) RemoveClient(client *WsClient) {
	delete(l.Clients, client)

	if err := l.Lobby.DisconnectPlayer(client.User); err != nil {
		fmt.Println(err)
	}
	if len(l.Clients) == 0 {
		l.hub.removeLobby(l.Id)
		return
	}
	event := dto.ConnectionEvent{
		IsConnected: false,
		User:        domain.User{ID: client.User.ID, Username: client.User.Username},
	}
	l.connectionsCh <- &event
	l.sendUpdates()
}

func (l *Lobby) MoveSlot(client *WsClient, from int, to int) {
	if client.User.ID != l.Lobby.OwnerId {
		fmt.Println("Only owner can move slot")
		return
	}
	l.Lobby.MoveSlot(from, to)
	l.sendUpdates()
}

// Старт игры
func (l *Lobby) StartGame(cl *WsClient) {
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

func (l *Lobby) sendUpdates() {
	event := dto.UpdateEvent{
		Type:  dto.EventUpdate,
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
