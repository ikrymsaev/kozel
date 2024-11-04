package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
)

type GameService struct {
	Game  *domain.Game
	Lobby *LobbyService
}

func NewGameService(lobby *LobbyService) GameService {
	game := domain.NewGame(lobby.Lobby)

	return GameService{
		Game:  &game,
		Lobby: lobby,
	}
}

func (g *GameService) Run() {
	g.Game.Start()

	fmt.Printf("Current Round FirstStepPlayer %v\n", g.Game.CurrentRound.FirstStepPlayer)
	fmt.Println("Game created")

	for client := range g.Lobby.Clients {
		client.gameStateCh <- &dto.GameStateEvent{
			Type: dto.EGameStateEvent,
			Game: g.Game,
		}
	}
}
