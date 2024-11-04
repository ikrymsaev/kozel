package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
)

type GameService struct {
	Game         *domain.Game
	LobbyService *LobbyService
}

func NewGameService(lobby *LobbyService) GameService {
	game := domain.NewGame(lobby.Lobby)

	return GameService{
		Game:         &game,
		LobbyService: lobby,
	}
}

func (g *GameService) PraiseTrump(cl *ClientService, trump domain.ESuit) {
	round := &g.Game.CurrentRound

	praiserId := round.Praiser.Id
	if praiserId != cl.User.ID {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "You are not praiser",
		}
		return
	}
	if round.Trump != nil {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "Trump is already set",
		}
		return
	}
	round.SetTrump(&trump)
	g.Game.SetStage(domain.StagePlayerStep)
	for client := range g.LobbyService.Clients {
		client.trumpCh <- &dto.NewTrumpEvent{
			Type:  dto.EventNewTrump,
			Trump: trump,
		}
		client.stageCh <- &dto.StageChangeEvent{
			Type:  dto.EventStageChange,
			Stage: domain.StagePlayerStep,
		}
	}
}

func (g *GameService) Run() {
	g.Game.Start()

	fmt.Printf("Current Round FirstStepPlayer %v\n", g.Game.CurrentRound.FirstStepPlayer)
	fmt.Println("Game created")

	for client := range g.LobbyService.Clients {
		client.gameStateCh <- &dto.GameStateEvent{
			Type: dto.EventGameState,
			Game: g.Game,
		}
	}
}
