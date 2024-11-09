package domain

import "fmt"

// Результат кона
type StakeResult struct {
	Winner *Player
	Bribe  []*Card
}

type Stake struct {
	Table       []*Card // Карты на кону
	round       *Round  // Ссылка на раунд
	CurrentStep *Player // Текущий игрок
}

func NewStake(round *Round, firstPlayer *Player) Stake {
	return Stake{
		Table:       []*Card{},
		round:       round,
		CurrentStep: firstPlayer,
	}
}

func (s *Stake) Start() StakeResult {
	println("Starting stake...")

	for i := 1; i <= 4; i++ {
		s.Turn()
	}

	winCard := GetWinCard(s.Table)

	return StakeResult{
		Winner: winCard.Owner,
		Bribe:  s.Table,
	}
}

func (s *Stake) SetPlayerTurn(player *Player) {
	s.CurrentStep = player
}

func (s *Stake) CalcResult() StakeResult {
	println("Getting stake result...")
	fmt.Printf("Table: %v\n", s.Table)
	winCard := GetWinCard(s.Table)
	return StakeResult{
		Winner: winCard.Owner,
		Bribe:  s.Table,
	}
}

// Переключатель текущего хода
func (s *Stake) Turn() {
	if s.CurrentStep.Position == 1 {
		s.CurrentStep = &s.round.Game.Teams[1].A
	} else if s.CurrentStep.Position == 2 {
		s.CurrentStep = &s.round.Game.Teams[0].B
	} else if s.CurrentStep.Position == 3 {
		s.CurrentStep = &s.round.Game.Teams[1].B
	} else if s.CurrentStep.Position == 4 {
		s.CurrentStep = &s.round.Game.Teams[0].A
	}
}

func (s *Stake) IsCompleted() bool {
	return len(s.Table) == 4
}

func (s *Stake) GetStakeSuit() *ESuit {
	if len(s.Table) == 0 {
		return nil
	}
	if s.Table[0].CardType.Type == Jack {
		return s.round.Trump
	}

	return &s.Table[0].CardSuit.Suit
}

func (s *Stake) PlayerAction(player *Player, cardId string) (*Card, error) {
	actionCard, err := player.PlayerAction(cardId)
	if err != nil {
		return actionCard, err
	}
	s.Table = append(s.Table, actionCard)
	actionCard.SetUsed()

	return actionCard, nil
}

// Действие игрока
func (s *Stake) BotAction(bot *Player) *Card {
	actionCard := bot.Action(s)
	if actionCard == nil {
		return nil
	}
	s.Table = append(s.Table, actionCard)
	actionCard.SetUsed()

	return actionCard
}
