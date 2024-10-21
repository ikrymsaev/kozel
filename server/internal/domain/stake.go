package domain

// Результат кона
type StakeResult struct {
	winner *Player
	bribe  []*Card
}

type Stake struct {
	cards         []*Card // Карты на кону
	round         *Round  // Ссылка на раунд
	currentPlayer *Player // Текущий игрок
}

func NewStake(round *Round, firstPlayer *Player) Stake {
	return Stake{
		cards:         []*Card{},
		round:         round,
		currentPlayer: firstPlayer,
	}
}

func (s *Stake) Start() StakeResult {
	println("Starting stake...")

	for i := 1; i <= 4; i++ {
		s.action()
		s.turn()
	}

	winCard := GetWinCard(&s.cards)

	return StakeResult{
		winner: winCard.Owner,
		bribe:  s.cards,
	}
}

// Переключатель текущего хода
func (s *Stake) turn() {
	if s.currentPlayer.Position == 1 {
		s.currentPlayer = &s.round.game.teams[1].A
	} else if s.currentPlayer.Position == 2 {
		s.currentPlayer = &s.round.game.teams[0].B
	} else if s.currentPlayer.Position == 3 {
		s.currentPlayer = &s.round.game.teams[1].B
	} else if s.currentPlayer.Position == 4 {
		s.currentPlayer = &s.round.game.teams[0].A
	}
}

func (s *Stake) GetStakeSuit() *ESuit {
	if len(s.cards) == 0 {
		return nil
	}
	if s.cards[0].CardType.Type == Jack {
		return s.round.trump
	}

	return &s.cards[0].CardSuit.Suit
}

// Действие игрока
func (s *Stake) action() {
	actionCard := s.currentPlayer.Action(s)
	s.cards = append(s.cards, actionCard)
	actionCard.SetUsed()
}
