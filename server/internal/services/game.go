package services

import (
	"fmt"
	"go-kozel/internal/domain"
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

func (s *GameService) Run() {
	s.Game.Start()

	fmt.Printf("Run FirstStepPlayer %v\n", s.Game.GetCurrentRound().FirstStepPlayer)
}
