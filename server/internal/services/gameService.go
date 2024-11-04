package services

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
	"time"
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

func (g *GameService) PraiseTrump(cl *ClientService, trump *domain.ESuit) {
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
	g.setTrump(trump)
}

func (g *GameService) setTrump(trump *domain.ESuit) {
	round := &g.Game.CurrentRound
	round.SetTrump(trump)
	g.Game.SetStage(domain.StagePlayerStep)
	for client := range g.LobbyService.Clients {
		client.trumpCh <- &dto.NewTrumpEvent{
			Type:  dto.EventNewTrump,
			Trump: *trump,
		}
		client.stageCh <- &dto.StageChangeEvent{
			Type:  dto.EventStageChange,
			Stage: domain.StagePlayerStep,
		}
	}
}

func (g *GameService) Run() {
	g.Game.Start()
	round := &g.Game.CurrentRound
	fmt.Printf("Current Round FirstStepPlayer %v\n", round.FirstStepPlayer)
	fmt.Println("Game created")

	for client := range g.LobbyService.Clients {
		client.gameStateCh <- &dto.GameStateEvent{
			Type: dto.EventGameState,
			Game: g.Game,
		}
	}

	if round.Praiser.User == nil {
		fmt.Printf("Bot praising trump \n")
		time.Sleep(3 * time.Second)
		trump := round.Praiser.PraiseTrump()
		g.setTrump(trump)
	}

}
