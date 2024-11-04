package dto

import (
	"fmt"
	"go-kozel/internal/domain"
)

type GameStateModel struct {
	Players [4]PlayerStateModel `json:"players"`
	Round   RoundStateModel     `json:"round"`
	Score   [2]byte             `json:"score"`
}

// Модель раунда
type RoundStateModel struct {
	FirstStepPlayerId string              `json:"firstStepPlayerId"`
	PraiserId         string              `json:"praiserId"`
	Trump             string              `json:"trump"`
	Bribes            [2][]CardStateModel `json:"bribes"`
}

func GetBribesStateModel(bribes [2][]*domain.Card) [2][]CardStateModel {
	bribesState := [2][]CardStateModel{}
	for i := 0; i < 2; i++ {
		bribesState[i] = []CardStateModel{}
		for _, card := range bribes[i] {
			bribesState[i] = append(bribesState[i], GetCardStateModel(card))
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

	fmt.Printf("firstStepPlayerId: %v\n", firstStepPlayerId)
	fmt.Printf("praiserId: %v\n", praiserId)
	fmt.Printf("trump: %v\n", trump)

	return RoundStateModel{
		FirstStepPlayerId: firstStepPlayerId,
		PraiserId:         praiserId,
		Trump:             trump,
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
}

func GetPlayerStateModel(player *domain.Player) PlayerStateModel {
	if player == nil {
		return PlayerStateModel{}
	}
	hand := []CardStateModel{}
	for _, card := range player.Hand {
		hand = append(hand, GetCardStateModel(card))
	}

	return PlayerStateModel{
		Id:       player.Id,
		Name:     player.Name,
		Hand:     hand,
		Position: player.Position,
		User:     player.User,
	}
}

// Модель карты
type CardStateModel struct {
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
		ImageUri: card.ImageUri,
		CardType: card.CardType.Name,
		CardSuit: card.CardSuit.Name,
		IsTrump:  card.IsTrump,
	}
}

func NewGameStateModel(game *domain.Game) GameStateModel {
	players := [4]PlayerStateModel{}
	for index, player := range game.GetPlayers() {
		players[index] = GetPlayerStateModel(&player)
	}

	return GameStateModel{
		Players: players,
		Round:   GetRoundStateModel(game.CurrentRound),
		Score:   game.Score,
	}
}
