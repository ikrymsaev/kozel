package dto

import (
	"fmt"
	"go-kozel/internal/domain"
)

type GameStateModel struct {
	Players [4]PlayerStateModel `json:"players"`
	Round   RoundStateModel     `json:"round"`
	Score   [2]byte             `json:"score"`
	Stage   domain.EStage       `json:"stage"`
}

// Модель раунда
type RoundStateModel struct {
	FirstStepPlayerId string  `json:"firstStepPlayerId"`
	PraiserId         string  `json:"praiserId"`
	TurnPlayerId      string  `json:"turnPlayerId"`
	Trump             string  `json:"trump"`
	Bribes            [2]byte `json:"bribes"`
}

func GetBribesStateModel(bribes [2][]*domain.Card) [2]byte {
	bribesState := [2]byte{}
	for i := 0; i < 2; i++ {
		for _, card := range bribes[i] {
			bribesState[i] += card.CardType.Score
		}
	}
	return bribesState
}

func GetRoundStateModel(round domain.Round) RoundStateModel {
	var firstStepPlayerId string
	if round.FirstStepPlayer != nil {
		firstStepPlayerId = round.FirstStepPlayer.Id
	}

	var praiserId string
	if round.Praiser != nil {
		praiserId = round.Praiser.Id
	}
	trump := ""
	if round.Trump != nil {
		trump = round.Trump.String()
	}
	turnPlayerId := ""

	if round.CurrentStake != nil && round.CurrentStake.CurrentStep != nil {
		turnPlayerId = round.CurrentStake.CurrentStep.Id
	}

	fmt.Printf("firstStepPlayerId: %v\n", firstStepPlayerId)
	fmt.Printf("praiserId: %v\n", praiserId)
	fmt.Printf("trump: %v\n", trump)

	return RoundStateModel{
		FirstStepPlayerId: firstStepPlayerId,
		PraiserId:         praiserId,
		Trump:             trump,
		TurnPlayerId:      turnPlayerId,
		Bribes:            GetBribesStateModel(round.Bribes),
	}
}

// Модель игрока
type PlayerStateModel struct {
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	Hand     []CardStateModel `json:"hand"`
	Position byte             `json:"position"`
	User     *domain.User     `json:"user"`
	Team     byte             `json:"team"`
}

func GetPlayerStateModel(player *domain.Player) PlayerStateModel {
	if player == nil {
		return PlayerStateModel{}
	}
	hand := []CardStateModel{}
	for _, card := range player.Hand {
		hand = append(hand, GetCardStateModel(card))
	}
	team := byte(0)
	if player.Team != nil {
		team = player.Team.Id
	}

	return PlayerStateModel{
		Id:       player.Id,
		Name:     player.Name,
		Hand:     hand,
		Position: player.Position,
		User:     player.User,
		Team:     team,
	}
}

// Модель карты
type CardStateModel struct {
	Id       string `json:"id"`
	IsHidden bool   `json:"isHidden"`
	ImageUri string `json:"imageUri"`
	CardType string `json:"type"`
	CardSuit string `json:"suit"`
	IsTrump  bool   `json:"isTrump"`
}

func GetCardStateModel(card *domain.Card) CardStateModel {
	if card == nil {
		return CardStateModel{}
	}
	return CardStateModel{
		Id:       card.Id,
		IsHidden: card.IsUsed,
		ImageUri: card.ImageUri,
		CardType: card.CardType.Name,
		CardSuit: card.CardSuit.Name,
		IsTrump:  card.IsTrump,
	}
}

func NewGameStateModel(game *domain.Game) GameStateModel {
	players := [4]PlayerStateModel{}
	for index, player := range game.GetPlayers() {
		players[index] = GetPlayerStateModel(player)
	}

	return GameStateModel{
		Players: players,
		Round:   GetRoundStateModel(game.CurrentRound),
		Score:   game.Score,
		Stage:   game.Stage,
	}
}

type StakeResultModel struct {
	WinnerId   string `json:"winnerId"`
	BribeScore byte   `json:"bribeScore"`
}

func GetStakeResultModel(stakeResult *domain.StakeResult) StakeResultModel {
	var score byte
	for _, card := range stakeResult.Bribe {
		score += card.CardType.Score
	}

	return StakeResultModel{
		WinnerId:   stakeResult.Winner.Id,
		BribeScore: score,
	}
}

type RoundResultModel struct {
	WinnerTeam byte `json:"winnerTeam"`
	Score      byte `json:"score"`
}

func GetRoundResultModel(roundResult *domain.RoundResult) RoundResultModel {
	team := roundResult.WinTeam.Id
	return RoundResultModel{
		WinnerTeam: team,
		Score:      roundResult.Score,
	}
}
