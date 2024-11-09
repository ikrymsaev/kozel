package game

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/dto"
	"time"
)

type Game struct {
	Game         *domain.Game
	LobbyService *Lobby
}

func NewGameService(lobby *Lobby) Game {
	game := domain.NewGame(lobby.Lobby)

	return Game{
		Game:         &game,
		LobbyService: lobby,
	}
}

func (g *Game) MoveCard(cl *WsClient, cardId string) {
	if g.Game.Stage != domain.StagePlayerStep {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "Wrong stage",
		}
		return
	}

	stake := g.Game.CurrentRound.CurrentStake
	player := g.Game.GetPlayerByUser(cl.User)
	if player == nil {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "Player is NOT User", //!! Must be action by bot
		}
		return
	}
	if stake.CurrentStep.Id != player.Id {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: "Not your turn",
		}
		return
	}

	// Action by User
	actionCard, actionErr := stake.PlayerAction(player, cardId)
	if actionErr != nil {
		cl.errorCh <- &dto.ErrorEvent{
			Type:  dto.EventError,
			Error: actionErr.Error(),
		}
		return
	}

	for client := range g.LobbyService.Clients {
		client.cardActionCh <- &dto.CardActionEvent{
			Type:     dto.EventCardAction,
			Card:     actionCard,
			PlayerId: player.Id,
		}
	}

	g.NextTurn()
}

func (g *Game) FinishStake() {
	stake := g.Game.CurrentRound.CurrentStake

	stage := domain.StageCalculation
	g.Game.SetStage(stage)
	for client := range g.LobbyService.Clients {
		client.stageCh <- &dto.StageChangeEvent{
			Type:  dto.EventStageChange,
			Stage: stage,
		}
	}

	time.Sleep(3 * time.Second)

	round := &g.Game.CurrentRound
	result := stake.CalcResult()
	round.AddBribe(&result)

	for client := range g.LobbyService.Clients {
		client.stakeResultCh <- &dto.StakeResultEvent{
			Type:   dto.EventStakeResult,
			Result: &result,
		}
	}

	if round.IsCompleted() {
		g.NextRound()
		return
	}

	g.Game.SetStage(domain.StagePlayerStep)
	round.InitStake()
	winner := result.Winner
	round.CurrentStake.SetPlayerTurn(winner)

	for client := range g.LobbyService.Clients {
		client.stageCh <- &dto.StageChangeEvent{
			Type:  dto.EventStageChange,
			Stage: domain.StagePlayerStep,
		}
		client.playerStepCh <- &dto.ChangeStepEvent{
			Type:       dto.EventChangeStep,
			PlayerStep: winner,
		}
	}

	if winner.IsBot() {
		g.BotMoveCard(winner)
	}
}

func (g *Game) NextRound() {
	round := &g.Game.CurrentRound
	result := round.GetResult()

	fmt.Printf("Round result %v\n", result)

	g.Game.AddScoreToTeam(&result)

	for client := range g.LobbyService.Clients {
		client.roundResultCh <- &dto.RoundResultEvent{
			Type:   dto.EventRoundResult,
			Result: &result,
		}
	}

	time.Sleep(2 * time.Second)

	g.Run()
}

func (g *Game) NextTurn() {
	stake := g.Game.CurrentRound.CurrentStake

	if stake.IsCompleted() {
		g.FinishStake()
		return
	}

	stake.Turn()
	nextStepPlayer := stake.CurrentStep

	for client := range g.LobbyService.Clients {
		client.playerStepCh <- &dto.ChangeStepEvent{
			Type:       dto.EventChangeStep,
			PlayerStep: nextStepPlayer,
		}
	}

	if nextStepPlayer.IsBot() {
		g.BotMoveCard(nextStepPlayer)
	}
}

func (g *Game) BotMoveCard(bot *domain.Player) {
	stake := g.Game.CurrentRound.CurrentStake
	fmt.Printf("Bot action player: %v\n", bot)
	// Action by Bot
	time.Sleep(2 * time.Second)
	fmt.Println("Bot after sleep")
	actionCard := stake.BotAction(bot)
	fmt.Printf("Bot selected card: %v\n", actionCard)

	for client := range g.LobbyService.Clients {
		client.cardActionCh <- &dto.CardActionEvent{
			Type:     dto.EventCardAction,
			Card:     actionCard,
			PlayerId: bot.Id,
		}
	}

	g.NextTurn()
}

func (g *Game) PraiseTrump(cl *WsClient, trump *domain.ESuit) {
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

	stepPlayer := round.CurrentStake.CurrentStep
	if stepPlayer.IsBot() {
		g.BotMoveCard(stepPlayer)
	}
}

func (g *Game) setTrump(trump *domain.ESuit) {
	round := &g.Game.CurrentRound
	round.SetTrump(trump)
	g.Game.SetStage(domain.StagePlayerStep)
	round.InitStake()

	stepPlayer := round.CurrentStake.CurrentStep

	for client := range g.LobbyService.Clients {
		client.trumpCh <- &dto.NewTrumpEvent{
			Type:  dto.EventNewTrump,
			Trump: *trump,
		}
		client.stageCh <- &dto.StageChangeEvent{
			Type:  dto.EventStageChange,
			Stage: domain.StagePlayerStep,
		}
		client.playerStepCh <- &dto.ChangeStepEvent{
			Type:       dto.EventChangeStep,
			PlayerStep: stepPlayer,
		}
	}
}

func (g *Game) Run() {
	gameWinner := g.GetGameWinner()
	if gameWinner != nil {
		for client := range g.LobbyService.Clients {
			client.stageCh <- &dto.StageChangeEvent{
				Type:  dto.EventStageChange,
				Stage: domain.StageGameOver,
			}
			client.gameOverCh <- &dto.GameOverEvent{
				Type:       dto.EventGameOver,
				WinnerTeam: gameWinner,
			}
		}
		return
	}

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

	if round.Praiser.IsBot() {
		fmt.Printf("Bot praising trump \n")
		time.Sleep(2 * time.Second)
		trump := round.Praiser.PraiseTrump()
		g.setTrump(trump)

		if round.FirstStepPlayer.IsBot() {
			fmt.Printf("Bot move card \n")
			time.Sleep(2 * time.Second)
			g.BotMoveCard(round.FirstStepPlayer)
		}
	}
}

func (g *Game) GetGameWinner() *domain.Team {
	score := g.Game.Score

	if score[0] >= 12 {
		return &g.Game.Teams[0]
	}

	if score[1] >= 12 {
		return &g.Game.Teams[1]
	}

	return nil
}
